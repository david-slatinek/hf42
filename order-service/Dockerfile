FROM openjdk:19-alpine

WORKDIR /app

COPY .env /app/.env
COPY /build/libs/*.jar /app/app.jar

CMD ["java", "-jar", "app.jar"]
