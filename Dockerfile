FROM teemow/arch

RUN pacman -Sy
RUN pacman -Syu --noconfirm
RUN pacman -Sy --noconfirm supercollider jack2
RUN pacman -Sy --noconfirm mplayer 

ADD startup.sh startup.sh

CMD ./startup.sh
