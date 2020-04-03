#!/bin/sh

gcloud components install cloud-spanner-emulator
gcloud beta emulators spanner start &

gcloud config configurations create emulator
gcloud config set auth/disable_credentials true
gcloud config set project $SPANNER_PROJECT_ID
gcloud config set api_endpoint_overrides/spanner http://localhost:9020/

gcloud spanner instances create $SPANNER_INSTANCE_ID --project=$SPANNER_PROJECT_ID --config=regional-us-central1 --description="" --nodes=1
gcloud spanner databases create $SPANNER_DATABASE_ID --project=$SPANNER_PROJECT_ID --instance=$SPANNER_INSTANCE_ID

gcloud spanner databases ddl update $SPANNER_DATABASE_ID --project=$SPANNER_PROJECT_ID --instance=$SPANNER_INSTANCE_ID --ddl="`cat $DDL`"
