from flask_restful import reqparse
import flask  
server=flask.Flask(__name__) 
from flask import request,jsonify

@server.route('/api',methods=['get'])

def reg():
    #parser = reqparse.RequestParser()
    #parser.add_argument('a', type=int, help='invalid argument')
    #parser.add_argument('b', type=str, help='invalid argument')
    #di= parser.parse_args()

    try:
        dit=request.args
        if len(dit)>2:
            return jsonify(dict(err_code=21,err_msg='empry or wrong params',reference=''))
        a = request.args.get('a', '')
        b = request.args.get('b', '')
        if not a or not b:
            return jsonify(dict(err_code=21,err_msg='empry or wrong params',reference=''))
        try:
            a = int(a)
            b=str(b)
        except ValueError:
            return jsonify(dict(err_code=21,err_msg='empry or wrong params',reference=''))
        return jsonify(dict(err_code=0,err_msg='success',reference='NO.{0} is {1}'.format(str(a),b)))
    except Exception as e:
        return jsonify(dict(err_code=11,err_msg='system error',reference=''))
    return jsonify(dict(refenence=a))
server.run(port=9999, debug=True)
#服务创建好，并指定好路径，接口内容编写完后，需要启动服务，启动的时候需要指定端口