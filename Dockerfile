FROM nginx:alpine

# Remove default index
RUN rm -rf /usr/share/nginx/html/*

# Copy our app
COPY signup.html /usr/share/nginx/html/index.html
