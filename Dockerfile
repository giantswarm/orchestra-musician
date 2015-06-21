FROM teemow/arch

RUN pacman -Sy
RUN pacman -Syu --noconfirm
RUN pacman -Sy --noconfirm supercollider jack2
RUN pacman -Sy --noconfirm mplayer 

COPY leiningen-1:2.5.1-1-any.pkg.tar.xz leiningen-1:2.5.1-1-any.pkg.tar.xz
RUN pacman --noconfirm -U leiningen-1:2.5.1-1-any.pkg.tar.xz
COPY . /musician 

WORKDIR /musician
RUN lein deps

CMD ./startup.sh
