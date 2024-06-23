import { Body, Controller, Post, Get } from '@nestjs/common';
import { AppService } from './app.service';
import { CreateUserInput } from 'src/shared/types/create-user.input';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  getData() {
    return this.appService.getData();
  }

  @Post('/users')
  createuser(@Body() createUserInput: CreateUserInput) {
    console.log('AppController');
    this.appService.createUser(createUserInput);
  }

  @Get()
  getHello(): string {
    return this.appService.getHello();
  }
}
