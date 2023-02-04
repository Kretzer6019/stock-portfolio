FROM node:18-alpine

# set the working direction
WORKDIR /app/frontend

# install app dependencies
COPY package.json ./
COPY package-lock.json ./

# Download packages
RUN yarn install

# Bundle app source
COPY . ./

# Run build
RUN yarn build

# start app
CMD ["yarn", "dev"]
