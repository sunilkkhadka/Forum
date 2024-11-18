export interface SuccessResponse<T = unknown> {
  response: {
    status: string;
    message: string;
    data: T;
  };
}
