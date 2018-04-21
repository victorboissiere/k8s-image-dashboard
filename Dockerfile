FROM mhart/alpine-node:8.11.1

WORKDIR /home/app

COPY package.json ./
COPY yarn.lock ./
RUN yarn install

COPY ./ ./

CMD ["yarn", "start"]

