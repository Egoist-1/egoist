.PHONY:webook
kratos:
#make kratos project=?
	@cd webook && kratos new $(project_name) -r https://gitee.com/go-kratos/kratos-layout.git
buf:
generate:
	@go generate ./...
	@go mod tidy
create:
	@./create.sh $(name)
wire:
	@cd webook/_bff && wire