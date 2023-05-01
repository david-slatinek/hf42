#FROM node:20 as build-deps
#WORKDIR /app
#COPY package.json package-lock.json yarn.lock ./
#RUN yarn
#COPY . ./
#RUN yarn build
#
#FROM nginx:1.23.4-alpine
#COPY --from=build-deps /app/build /usr/share/nginx/html
#EXPOSE 3000
#CMD ["nginx", "-g", "daemon off;"]

#FROM node:14-alpine
#WORKDIR /app
#COPY package.json ./
#RUN npm install
#COPY . .
#CMD ["npm", "start"]

FROM node:20-alpine3.16
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile
COPY . .
RUN yarn build
ENV NODE_ENV=production
CMD ["yarn", "start"]