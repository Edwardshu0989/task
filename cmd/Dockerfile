FROM scratch

WORKDIR /usr/src/test-server

COPY . .

RUN chmod +x ./awesomeProject

EXPOSE 8082

ENTRYPOINT ["./awesomeProject"]
