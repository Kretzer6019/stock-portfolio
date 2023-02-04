FROM node:lts-alpine

# set the working direction
WORKDIR /app/frontend

# install app dependencies
COPY package.json ./
COPY package-lock.json ./

# Fix permissions for packages
# RUN npm config set unsafe-perm true

RUN npm install
RUN npm install -g serve

# Bundle app source
COPY . ./

RUN chown -R node:node /app/frontend/node_modules

# start app
CMD ["npm", "run", "dev"]
