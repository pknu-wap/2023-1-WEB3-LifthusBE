import { Injectable, CanActivate, ExecutionContext } from '@nestjs/common';
import { Observable } from 'rxjs';

// only signed user can access with req.uid
@Injectable()
export class UserGuard implements CanActivate {
  canActivate(
    context: ExecutionContext,
  ): boolean | Promise<boolean> | Observable<boolean> {
    const uid = context.switchToHttp().getRequest().uid;
    // undefined means not signed user so block.
    // and uid 0 can't exist. autoincrement starts from 1
    if (!uid) {
      return false;
    }
    return true;
  }
}