export = Kiko;

declare class Kiko {
  constructor();
  fetch(input: RequestInfo, init?: RequestInit): Promise<Response>;
}
