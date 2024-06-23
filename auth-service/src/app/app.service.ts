import { Injectable } from '@nestjs/common';
import { CreateUserEvent } from 'src/modules/create-user.event';

@Injectable()
export class AppService {
  getHello(): string {
    return 'Hello World!';
  }

  getData(): { message: string } {
    return { message: 'auth-service' };
  }

  handleUserCreated(data: CreateUserEvent) {
    console.log('handleUserCreated - AUTH-SERVICE', data);
  }
}
