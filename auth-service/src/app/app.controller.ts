import { Controller, Get } from '@nestjs/common';
import { AppService } from './app.service';
import { EventPattern } from '@nestjs/microservices';
import { CreateUserEvent } from 'src/modules/create-user.event';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  getData() {
    return this.appService.getData();
  }

  @EventPattern('user_created')
  handleuserCreated(data: CreateUserEvent) {
    this.appService.handleUserCreated(data);
  }

  @Get()
  getHello(): string {
    return this.appService.getHello();
  }
}
