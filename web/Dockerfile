FROM node:19 as node
WORKDIR /app
COPY . .
RUN npm install && npm run build

FROM nginx:1.23-alpine
COPY --from=node ./app/build /usr/share/nginx/html
CMD nginx -g "daemon off;"
