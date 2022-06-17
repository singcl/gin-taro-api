/* global Promise */
import Md5Con from '../../lib/authorization/md5.min.js';
import generateAuthorization from './../generateAuthorization.js';

//
export default class Kiko {
  businessKey = 'admin';
  businessSecret = '12878dd962115106db6d';

  /**
   * @param {{businessKey: string, businessSecret: string}} params
   */
  constructor(params) {
    const businessKey = params && params.businessKey;
    const businessSecret = params && params.businessSecret;
    if (typeof businessKey !== 'undefined') this.businessKey = businessKey;
    if (typeof businessSecret !== 'undefined')
      this.businessSecret = businessSecret;
  }
  /**
   *
   * @param {RequestInfo} input
   * @param {RequestInit} init
   * @returns
   */
  async fetch(input, init = {}) {
    const method = init.method || 'GET';
    const url = typeof input === 'string' ? input : input && input.url;
    const body = init.body;
    //
    let password = body && body.password;
    password = password ? Md5Con.md5(password) : password;

    let bodyR = body && Object.assign({}, body, { password });
    //
    const authorizationData = generateAuthorization({
      path: url,
      method,
      params: bodyR,
      businessKey: this.businessKey,
      businessSecret: this.businessSecret,
    });
    //
    const token = localStorage.getItem(Kiko.getTokenName());
    /**
     * @type {RequestInit}
     */
    const initAuth = Object.assign({}, init, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8',
        Authorization: authorizationData.authorization,
        'Authorization-Date': authorizationData.date,
        Token: token,
      },
      body: bodyR ? new URLSearchParams(bodyR) : undefined,
    });
    //
    const response = await fetch(input, initAuth);
    const status = response.status;
    const resBody = await response.json();
    const ok = response.ok;
    if (ok && status == 200) {
      return resBody;
    } else {
      return Promise.reject(resBody);
    }
  }

  static getTokenName() {
    return '_login_token_';
  }
}
