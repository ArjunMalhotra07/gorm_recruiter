.PHONY: job
job:
	@docker build -t arjunmalhotra07/job-portal-service:latest .

.PHONY: email
email:
	@docker build -t arjunmalhotra07/email-service:latest ../email-service

.PHONY: server
server:
	@docker compose -f docker-compose.yml up --build