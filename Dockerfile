FROM teemow/arch

RUN pacman -Sy
RUN pacman -Syu --noconfirm
RUN pacman -Sy --noconfirm vlc jack2
# WARNING, horrible hack to enable vlc to run as root!
RUN sed -i 's/geteuid/getppid/' /usr/bin/vlc

RUN pacman -Sy --noconfirm mplayer

WORKDIR /musician
COPY . /musician
COPY startup.sh startup.sh
ENTRYPOINT ["./startup.sh"]
