import { NestFactory } from '@nestjs/core';
import { AppModule } from './app/app.module';
import { Transport, MicroserviceOptions } from '@nestjs/microservices';
import { Logger } from '@nestjs/common';

const logger = new Logger();

async function bootstrap() {
  const app = await NestFactory.createMicroservice<MicroserviceOptions>(
    AppModule,
    {
      transport: Transport.KAFKA,
      options: {
        client: {
          clientId: 'auth-service',
          brokers: ['localhost:9092'],
        },
        consumer: {
          groupId: 'auth-service',
        },
      },
    },
  );
  logger.log('Auth-service is listening');
  await app.listen();
}
bootstrap();
