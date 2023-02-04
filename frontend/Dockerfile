FROM node:18-alpine

# set the working direction
WORKDIR /app/frontend

# install app dependencies
COPY package.json ./
COPY package-lock.json ./

# Download packages
RUN npm install

# Bundle app source
COPY . ./

# Run build
RUN npm run build

# start app
CMD ["npm", "run", "dev"]