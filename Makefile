.PHONY: yo
yo: 
	yo $(SPANNER_PROJECT_ID) $(SPANNER_INSTANCE_ID) $(SPANNER_DATABASE_ID) -o models

.PHONY: db/reset
db/reset: 
	wrench --project $(SPANNER_PROJECT_ID) --instance $(SPANNER_INSTANCE_ID) --database $(SPANNER_DATABASE_ID) drop
	wrench --project $(SPANNER_PROJECT_ID) --instance $(SPANNER_INSTANCE_ID) --database $(SPANNER_DATABASE_ID) --directory db create

.PHONY: test
test: 
	go test ./...
