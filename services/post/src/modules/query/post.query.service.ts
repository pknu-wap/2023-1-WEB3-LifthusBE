import { Injectable } from '@nestjs/common';
import { Response } from 'express';

@Injectable()
export class PostQueryService {
  getHello(): string {
    return 'Hello World!';
  }
}
