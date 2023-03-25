FROM openjdk:19-alpine

WORKDIR /app

COPY env.properties /app/env.properties
COPY /build/libs/user-service-0.0.1-SNAPSHOT.jar /app/app.jar

ENV debug=false

CMD ["java", "-jar", "app.jar"]
