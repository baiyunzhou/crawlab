FROM node:14 AS build

ADD . /app
WORKDIR /app
RUN rm /app/.npmrc

# install frontend
ENV SASS_BINARY_SITE=https://npm.taobao.org/mirrors/node-sass/
RUN npm config set registry "http://registry.npm.taobao.org"
RUN npm i -g pnpm@7
RUN pnpm install
RUN pnpm run build

FROM nginx:1.27.2

# copy files
COPY --from=build /app/dist /app/dist
