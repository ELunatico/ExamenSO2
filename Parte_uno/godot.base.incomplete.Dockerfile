FROM ubuntu:24.04

# 1. Instala todas las dependencias necesarias con configuración completa
RUN apt-get update && apt-get install -y \
    wget \
    unzip \
    fontconfig \
    libfontconfig1 \
    libfreetype6 \
    fonts-dejavu-core \
    fonts-liberation2 \
    fonts-noto-core \
    libx11-6 \
    libxcursor1 \
    libxinerama1 \
    libxrandr2 \
    libxi6 \
    libasound2t64 \
    && rm -rf /var/lib/apt/lists/*

# 2. Configuración explícita del entorno de fuentes
RUN mkdir -p /usr/share/fonts/truetype \
    && ln -s /usr/share/fonts/truetype/dejavu /usr/share/fonts/truetype/DejaVuSans.ttf \
    && fc-cache -fv

# 3. Descarga e instala Godot
RUN wget https://github.com/godotengine/godot/releases/download/4.4-stable/Godot_v4.4-stable_linux.x86_64.zip -O /tmp/godot.zip \
    && unzip /tmp/godot.zip -d /tmp \
    && mv /tmp/Godot_v4.4-stable_linux.x86_64 /usr/local/bin/godot \
    && chmod +x /usr/local/bin/godot \
    && rm /tmp/godot.zip

# 4. Variables de entorno para Godot
ENV GODOT_FONT_CONFIG_PATH=/etc/fonts
ENV GODOT_FONT_CACHE_PATH=/var/cache/fontconfig

# 5. Comando por defecto
CMD ["godot", "--headless"]