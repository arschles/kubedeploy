FROM alpine:3.3

ADD kubedeploy /bin/kubedeploy

CMD ["/bin/kubedeploy"]
