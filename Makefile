#################################################################################
PPROF=go tool pprof
MYSQL=sudo mysql -h$(DB_HOST) -P$(DB_PORT) -u$(DB_USER) -p$(DB_PASS) $(DB_NAME)
SLACKCAT=slackcat --channel $(SLACKCAT_CNL)
WHEN:=$(shell date +%H:%M:%S)
#################################################################################

.PHONY: pull
pull:
	cd $(GIT_ROOT) && \
		git pull

.PHONY: restart
restart: restart-nginx restart-mysql restart-app

.PHONY: restart-app
restart-app:
	cd ~/isuumo/webapp/go && make all
	sudo systemctl restart isuumo.go.service

.PHONY: restart-nginx
restart-nginx:
	sudo rm -f /var/log/nginx/access.log
	sudo nginx -t
	sudo systemctl reload nginx

.PHONY: restart-mysql
restart-mysql:
	sudo rm -f /var/log/mysql/mysql-slow.log
	sudo systemctl restart mysql

.PHONY: app-log
app-log:
	sudo journalctl -u isuumo.go.service

.PHONY: alp
alp:
