FROM alpine:3.6

RUN adduser -D kube-scheduler-operator
USER kube-scheduler-operator

ADD tmp/_output/bin/kube-scheduler-operator /usr/local/bin/kube-scheduler-operator
