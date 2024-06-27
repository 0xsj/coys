import { Inject, Injectable } from '@nestjs/common';
import { ClientProxy } from '@nestjs/microservices';
import { CreateUserInput } from 'src/shared/types/create-user.input';
import { CreatedUserEvent } from 'src/shared/events/create-user.event';

@Injectable()
export class AppService {
  private readonly users: any[] = [];

  constructor(
    @Inject('AUTH-SERVICE') private readonly communicationClient: ClientProxy,
  ) {}

  getData(): { message: string } {
    return { message: 'gateway backend' };
  }

  createUser(createUserInput: CreateUserInput) {
    console.log('create user()');
    this.users.push(createUserInput);
    this.communicationClient.emit(
      'user-created',
      new CreatedUserEvent(createUserInput.email),
    );
  }

  getHello(): string {
    return 'Hello World!';
  }
}
