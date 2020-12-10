FROM node:14 as build

WORKDIR /build

COPY ./pages ./pages
COPY ./public ./public
COPY ./components ./components
COPY ./package.json .
COPY ./yarn.lock .
COPY ./next.config.js .

RUN yarn install
RUN yarn build

EXPOSE 3000
CMD yarn start
