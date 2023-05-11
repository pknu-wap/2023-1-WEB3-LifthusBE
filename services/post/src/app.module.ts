import { MiddlewareConsumer, Module } from '@nestjs/common';
import { OpenapiController } from './openapi/openapi.controller';
import { JwtModule } from '@nestjs/jwt';
import { OpenapiService } from './openapi/openapi.service';
import { UidMiddleware } from './common/middlewares/uid.middleware';
import { QueryController } from './query/query.controller';
import { MutationController } from './mutation/mutation.controller';
import { QueryService } from './query/query.service';
import { MutationService } from './mutation/mutation.service';

@Module({
  imports: [
    JwtModule.register({
      secret: process.env.HUS_SECRET_KEY,
    }),
  ],
  controllers: [QueryController, MutationController, OpenapiController],
  providers: [QueryService, MutationService, OpenapiService],
})
export class AppModule {
  configure(consumer: MiddlewareConsumer) {
    consumer.apply(UidMiddleware).forRoutes('*');
  }
}
