
FROM node:23 AS build
WORKDIR /app


COPY package*.json ./
RUN npm install --legacy-peer-deps


COPY . .
RUN npm run build


FROM node:23-alpine

   
WORKDIR /app


COPY --from=build /app ./


EXPOSE 3000


CMD ["npm", "start"]
