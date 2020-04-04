FROM google/cloud-sdk:287.0.0-debian_component_based

RUN curl -O https://dl.google.com/go/go1.14.1.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.1.linux-amd64.tar.gz
RUN mv /usr/local/go/bin/* /bin

RUN gcloud components install cloud-spanner-emulator
