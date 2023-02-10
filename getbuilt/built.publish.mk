# this is a workaround because manifest-tool does not work for ecr. rather than
# pushing to a unique repo per architecture, ecr expects a single repo with each
# arch tagged. this alternate process, retags the build results and uses docker
# manifest rather than manifest-tool.

define built.repo.targets
built.img.release.publish.$(1).$(2).$(3):
	@$(INFO) docker tag $(1)/$(2)-$(3):$(VERSION) $(1)/$(2):$(VERSION)-$(3)
	@docker tag $(1)/$(2)-$(3):$(VERSION) $(1)/$(2):$(VERSION)-$(3) || $(FAIL)
	@$(OK) docker tag $(1)/$(2)-$(3):$(VERSION) $(1)/$(2):$(VERSION)-$(3)
	@$(INFO) docker push $(1)/$(2):$(VERSION)-$(3)
	@docker push $(1)/$(2):$(VERSION)-$(3) || $(FAIL)
	@$(OK) docker push $(1)/$(2):$(VERSION)-$(3)
	@$(INFO) docker manifest create --amend $(1)/$(2):$(VERSION) $(1)/$(2):$(VERSION)-$(3)
	@docker manifest create --amend $(1)/$(2):$(VERSION) $(1)/$(2):$(VERSION)-$(3) || $(FAIL)
	@$(OK) docker manifest create --amend $(1)/$(2):$(VERSION) $(1)/$(2):$(VERSION)-$(3)
	@$(INFO) docker manifest annotate --arch $(3) $(1)/$(2):$(VERSION) $(1)/$(2):$(VERSION)-$(3)
	@docker manifest annotate --arch $(3) $(1)/$(2):$(VERSION) $(1)/$(2):$(VERSION)-$(3) || $(FAIL)
	@$(OK) docker manifest annotate --arch $(3) $(1)/$(2):$(VERSION) $(1)/$(2):$(VERSION)-$(3)
built.img.release.publish: built.img.release.publish.$(1).$(2).$(3)
endef
$(foreach r,$(REGISTRY_ORGS), $(foreach i,$(IMAGES), $(foreach a,$(IMAGE_ARCHS),$(eval $(call built.repo.targets,$(r),$(i),$(a))))))

built.img.release.manifest.publish.%: built.img.release.publish
	@$(INFO) docker manifest push --purge $(DOCKER_REGISTRY)/$*:$(VERSION)
	@docker manifest push --purge $(DOCKER_REGISTRY)/$*:$(VERSION) || $(FAIL)
	@$(OK) docker manifest push --purge $(DOCKER_REGISTRY)/$*:$(VERSION)


# only publish images for main and release branches
ifneq ($(filter main release-%,$(BRANCH_NAME)),)
built.publish.artifacts: $(addprefix built.img.release.manifest.publish.,$(IMAGES))
endif

# publish all releasable artifacts
built.publish: version.isdirty
	@echo "$(BRANCH_NAME) $(PLATFORMS) $(IMAGES) $(REGISTRY_ORGS)"
	@$(MAKE) publish.init
	@$(MAKE) built.publish.artifacts
