# Use the official Flutter base image
FROM instrumentisto/flutter:3.13.5

# Set the environment variables for Flutter
ENV PATH="$PATH:/flutter/bin/cache/dart-sdk/bin:/flutter/bin"
ENV FLUTTER_ROOT=/flutter
ENV FLUTTER_STORAGE_BASE_URL=https://storage.flutter-io.cn

# Accept Android licenses non-interactively
RUN yes | flutter doctor --android-licenses

# Install Android SDK components
RUN flutter doctor --android-licenses

# Set the environment variables for Android
ENV ANDROID_HOME=/flutter/.pub-cache/bin/cache/android-sdk
ENV PATH="$PATH:$ANDROID_HOME/tools:$ANDROID_HOME/tools/bin:$ANDROID_HOME/platform-tools"

# Create a working directory for your Flutter app
WORKDIR /app

# Copy the Flutter app source code into the container
COPY . .