FROM scratch

ENV TZ Asia/Shanghai

WORKDIR /app
COPY ./main ./

CMD ["./goapis"]