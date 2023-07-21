##
## bulid web
##
FROM node:18-alpine as build
WORKDIR /app
COPY . .

# RUN npm install tyarn -g
# RUN tyarn
# RUN yarn build

RUN npm config set registry https://registry.npmmirror.com/
RUN npm i -g npm
RUN npm install
RUN npm run build
RUN npm run export


##
## deploy
##
FROM openresty/openresty:alpine
COPY --from=build /app/out /usr/local/openresty/nginx/html

EXPOSE 80 443
CMD ["/usr/local/openresty/bin/openresty", "-g", "daemon off;"]

