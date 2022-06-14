export default class Kiko {
  constructor() {
    //
  }
  /**
   *
   * @param {RequestInfo} input
   * @param {RequestInit} init
   * @returns
   */
  fetch(input, init) {
    /**
     * @type {RequestInit}
     */
    const authInit = {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8',
        Authorization: authorizationData.authorization,
        'Authorization-Date': authorizationData.date,
        Token: $.cookie('_login_token_'),
      },
    };
    return fetch(input, Object.assign(authInit, init));
  }
}
