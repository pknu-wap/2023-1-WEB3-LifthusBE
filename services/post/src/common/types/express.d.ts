// to make the file a module and avoid the TypeScript error
export {};

declare global {
  namespace Express {
    type Uid = number;
    type Exp = boolean;
    interface Request {
      uid?: Uid;
      exp?: Exp;
    }
  }
}
