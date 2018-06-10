# Copyright (c) 2018 Harry Roberts. All Rights Reserved.
# SPDX-License-Identifier: LGPL-3.0+

# Adapted from https://gist.githubusercontent.com/HarryR/d2373f421c39353cd462/raw/c3f726455bbbc03fe791dda0dabbf4f73b5d2ec9/restclient.py
# Except the command-line interface has been removed (so it doesn't depend on Plugin and Host components)

__all__ = ('RestClient',)

try:
    from urllib.parse import quote_plus
except ImportError:
    from urllib import quote_plus
import requests


class Resource(object):
    __slots__ = ('_api', '_url')

    def __init__(self, url=None, api=None):
        self._url = url
        self._api = api

    def __getattr__(self, name):
        if name[0] == '_':
            raise AttributeError
        return Resource(self._url + '/' + name, self._api)

    def __call__(self, name=None):
        if name is None:
            return self
        if name[0] == '_':
            raise AttributeError
        return Resource(self._url + '/' + quote_plus(name), self._api)

    def _do(self, method, kwargs):
        resp = self._api._request(method=method,
                                  url=self._url,
                                  params=kwargs)
        resp.raise_for_status()
        return resp

    def GET(self, **kwargs):
        return self._do('GET', kwargs)

    def POST(self, **kwargs):
        return self._do('POST', kwargs)

    def PUT(self, **kwargs):
        return self._do('PUT', kwargs)

    def DELETE(self, id=None, **kwargs):
        url = self._url
        if id is not None:
            url += "/" + str(id)
        resp = self._api._request(method='DELETE',
                                  url=url,
                                  data=kwargs)
        resp.raise_for_status()
        return resp


class RestClient(object):
    __slots__ = ('_base_url', '_session')
    def __init__(self, base_url):
        self._base_url = base_url.rstrip('/')
        self._session = requests.Session()

    def _request(self, **kwargs):
        req = requests.Request(**kwargs)
        # TODO: setup authentication on `req`
        return self._session.send(req.prepare())

    def __getattr__(self, name):        
        return Resource(self._base_url + '/' + name, self)
