FROM node:16.14.2 AS BUILDER
RUN npm install -g cnpm --registry=https://registry.npm.taobao.org
RUN mkdir /src
COPY . /src
WORKDIR /src
RUN cnpm install
RUN cnpm run build

FROM nginx:1.21.4
COPY --from=builder /src/dist /usr/share/nginx/html
COPY ./deploy/nginx.conf /etc/nginx/conf.d/default.conf
ENTRYPOINT ["nginx", "-g", "daemon off;"]
