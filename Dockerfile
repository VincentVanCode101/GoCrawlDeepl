ARG APP_BASE_IMAGE=golang:1.22.0

FROM ${APP_BASE_IMAGE} as dev

WORKDIR /usr/src/app

RUN curl https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb -o /tmp/google-chrome.deb

RUN apt update && apt install -y \                                                                                           
    /tmp/google-chrome.deb

    RUN apt update && apt install -y \
    xclip \
    sudo \
    && rm -rf /var/lib/apt/lists/*

COPY ./app/go.mod ./
COPY ./app/go.sum ./
RUN go mod download && go mod verify

COPY ./app .

RUN mkdir -p /usr/local/bin/app
RUN go build -v -o /usr/local/bin/app ./...

ARG USER_ID
ARG GROUP_ID
ARG USERNAME

RUN addgroup --system --gid ${GROUP_ID} ${USERNAME} \
    && adduser --system --uid ${USER_ID} --ingroup ${USERNAME} --home /home/${USERNAME} ${USERNAME}

RUN echo "${USERNAME} ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/${USERNAME}

USER ${USERNAME}

CMD [ "tail", "-f", "/dev/null" ]
