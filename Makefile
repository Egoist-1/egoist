.PHONY:webook
front:
	@cd webook_fe && npm run dev
back:
	@./setup.sh
kratos:
#make kratos project=?
	@cd webook && kratos new $(project_name) -r https://gitee.com/go-kratos/kratos-layout.git
buf:
mock:
docker:
	@cd webook && docker-compose up -d
create:
	@./create.sh $(name)