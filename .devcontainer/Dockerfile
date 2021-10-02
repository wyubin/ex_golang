FROM golang:1.16.6-alpine
RUN apk --no-cache add make gcc musl-dev git bash sudo

ARG USERNAME=yubin
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN adduser -D -u $USER_UID $USERNAME \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# RUN groupmod --gid $USER_GID $USERNAME \
#     && usermod --uid $USER_UID --gid $USER_GID $USERNAME \
#     && chown -R $USER_UID:$USER_GID /home/$USERNAME