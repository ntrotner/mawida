FROM node:20-alpine3.17
RUN apk add --no-cache git
COPY ./src/ui /src
WORKDIR /src
RUN npm install
RUN npm i -g vite
RUN npm run postinstall
ENTRYPOINT if [ "$app_env" = "prod" ] ; then npm run preview ; else npm run dev ; fi
