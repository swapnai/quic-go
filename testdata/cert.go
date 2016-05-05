package testdata

import (
	"crypto/tls"
	"crypto/x509"
)

// GetTLSConfig returns a tls config for quic.clemente.io
func GetTLSConfig() *tls.Config {
	return &tls.Config{
		Certificates: []tls.Certificate{GetCertificate()},
	}
}

// GetCertificate returns a certificate for quic.clemente.io
func GetCertificate() tls.Certificate {
	key, err := x509.ParsePKCS1PrivateKey(keyDer)
	if err != nil {
		panic(err)
	}
	return tls.Certificate{
		Certificate: [][]byte{certDer},
		PrivateKey:  key,
	}
}

var certDer = []byte("0\x82\x05\x040\x82\x03\xec\xa0\x03\x02\x01\x02\x02\x12\x03\xde)\r\xc6N\xfd\x11*Z\xd5\xf3\xce\xf5\x8c:F\x970\r\x06\t*\x86H\x86\xf7\r\x01\x01\v\x05\x000J1\v0\t\x06\x03U\x04\x06\x13\x02US1\x160\x14\x06\x03U\x04\n\x13\rLet's Encrypt1#0!\x06\x03U\x04\x03\x13\x1aLet's Encrypt Authority X30\x1e\x17\r160415160800Z\x17\r160714160800Z0\x1b1\x190\x17\x06\x03U\x04\x03\x13\x10quic.clemente.io0\x82\x01\"0\r\x06\t*\x86H\x86\xf7\r\x01\x01\x01\x05\x00\x03\x82\x01\x0f\x000\x82\x01\n\x02\x82\x01\x01\x00\xb7D\x11dAY0\x02\"u(HfCT\xb7\xc6E\x9e\xbfz\x80\x04\x14\xa3\xff`\xb1{\u02c8\v\xab\xfc\xa3/\xe7W\xb1c\xf4\x91e\xbd\xee\xd2\xdeBC\xfd\"f\xed\x01S\x17\x8d\xaf2C\xd5\x04m\xb1?\x88&\x8b\b\x02\xd1\xdc\x1e?\x80U=\xf3\x8f\xea\bL\xa7=\xa2\xc0NhY\xa4I_\xc8TV\x04\x86\x85.\x0e\x1a8\x12\x86v\v\xce\xdb%v{\xf3\xd3\xfb\xe4\x02\x8e\xd6\xf7\xe0\x98\x99E]W\xbe\xced\x10\xfaJcx\x1f\xa5t\x0f\xa0]\x8d\xb7\xa5\x96\xe6TT\x8e\xa3p!\xbd\x04\x9cR\x84I\xae\x88*[\x04<3:\u3e92<\u060c\x97\xe5:\xa9\x8a\xc1\x9b\xbdV~|\xc4k\x97\x8f.\u008d\u04db\xed\x02_\u33c6\u01c4\xac\xbd_\x9b\x8a\xc5\xfcn\xb8\x9cH\xb0Ug\xa5R\xc2i\u0082x\xad>\xb4\\\xcbl\u007f\xe4h0\xbd\x1e\x87\x8a\xd7\xd4yf\x95\x1d9\x9f\xfam\xf9e\t\x80n\xf4yr\xabz$o\xbf\x02\x03\x01\x00\x01\xa3\x82\x02\x110\x82\x02\r0\x0e\x06\x03U\x1d\x0f\x01\x01\xff\x04\x04\x03\x02\x05\xa00\x1d\x06\x03U\x1d%\x04\x160\x14\x06\b+\x06\x01\x05\x05\a\x03\x01\x06\b+\x06\x01\x05\x05\a\x03\x020\f\x06\x03U\x1d\x13\x01\x01\xff\x04\x020\x000\x1d\x06\x03U\x1d\x0e\x04\x16\x04\x14\x890\x12\x12\u0379\u68190\x0f\xac\x13\x98\x80\xd2?\xfc\x8f\xa60\x1f\x06\x03U\x1d#\x04\x180\x16\x80\x14\xa8Jjc\x04}\u077a\xe6\xd19\xb7\xa6Ee\xef\xf3\xa8\xec\xa10p\x06\b+\x06\x01\x05\x05\a\x01\x01\x04d0b0/\x06\b+\x06\x01\x05\x05\a0\x01\x86#http://ocsp.int-x3.letsencrypt.org/0/\x06\b+\x06\x01\x05\x05\a0\x02\x86#http://cert.int-x3.letsencrypt.org/0\x1b\x06\x03U\x1d\x11\x04\x140\x12\x82\x10quic.clemente.io0\x81\xfe\x06\x03U\x1d \x04\x81\xf60\x81\xf30\b\x06\x06g\x81\f\x01\x02\x010\x81\xe6\x06\v+\x06\x01\x04\x01\x82\xdf\x13\x01\x01\x010\x81\xd60&\x06\b+\x06\x01\x05\x05\a\x02\x01\x16\x1ahttp://cps.letsencrypt.org0\x81\xab\x06\b+\x06\x01\x05\x05\a\x02\x020\x81\x9e\f\x81\x9bThis Certificate may only be relied upon by Relying Parties and only in accordance with the Certificate Policy found at https://letsencrypt.org/repository/0\r\x06\t*\x86H\x86\xf7\r\x01\x01\v\x05\x00\x03\x82\x01\x01\x00\x1bo\xf3U\azq\xd6OUG\xa1K`\x81\x8b\xb5c\xf3\xb0\xa3Q\x05\xd9\xf8\xa7\x9aU\u0202\x96\xfc\xec\xcc^\xcc\xe1P\xd8j\xfeg\xf4\xab\x97\xe7v\x98\x1f>P\x1cPX\xc0Z\xef\x05Gf\xff\x81\xd9q\x96\xe2\x12\fM\xb2\xf2#\x1d\xee\xd7\xf7\xc8f\x11\ud5bcw(\x83\xc90\xbb\xba\x1aA\r\r\x1b\xa9\xc9\x19t\xa9bIEd\x8a\xdf\x19\u0154^\xb7\x95a_\xe2\x00\xaa\xfa\xf1Tk\xfb\xd2F\xa2l\xf2\xbdj\xdb\xe6q\x1d\x9d\xce\\G\x93='?\x89\xa5\x12\xecY\xf8\u9949<\x90\xbd\xc3\xf1\xb8\xbf&\xd1!\xc3%\u019353S7I\xf9Q\xb3RyY\xb3j\x81\xf6\rv\u007fY\x9a\xc4\x14\xa1\xf2\xd0\xe9\f\xf6W]\xf0\x8a\xad\x02\xediqlx\xc8\xd5\x18i\xc3\u0452\xbcw\x83\x9f\xb7\xb8'H@\x0f\xbd8\xb4v\x94\xac\xa2(]I\xa4\x91\xd0\x05i\xc9FS\xb8\xf7~ \xac\xba!\x94{YB\x93\u0469&J}E%")

var keyDer = []byte("0\x82\x04\xa5\x02\x01\x00\x02\x82\x01\x01\x00\xb7D\x11dAY0\x02\"u(HfCT\xb7\xc6E\x9e\xbfz\x80\x04\x14\xa3\xff`\xb1{\u02c8\v\xab\xfc\xa3/\xe7W\xb1c\xf4\x91e\xbd\xee\xd2\xdeBC\xfd\"f\xed\x01S\x17\x8d\xaf2C\xd5\x04m\xb1?\x88&\x8b\b\x02\xd1\xdc\x1e?\x80U=\xf3\x8f\xea\bL\xa7=\xa2\xc0NhY\xa4I_\xc8TV\x04\x86\x85.\x0e\x1a8\x12\x86v\v\xce\xdb%v{\xf3\xd3\xfb\xe4\x02\x8e\xd6\xf7\xe0\x98\x99E]W\xbe\xced\x10\xfaJcx\x1f\xa5t\x0f\xa0]\x8d\xb7\xa5\x96\xe6TT\x8e\xa3p!\xbd\x04\x9cR\x84I\xae\x88*[\x04<3:\u3e92<\u060c\x97\xe5:\xa9\x8a\xc1\x9b\xbdV~|\xc4k\x97\x8f.\u008d\u04db\xed\x02_\u33c6\u01c4\xac\xbd_\x9b\x8a\xc5\xfcn\xb8\x9cH\xb0Ug\xa5R\xc2i\u0082x\xad>\xb4\\\xcbl\u007f\xe4h0\xbd\x1e\x87\x8a\xd7\xd4yf\x95\x1d9\x9f\xfam\xf9e\t\x80n\xf4yr\xabz$o\xbf\x02\x03\x01\x00\x01\x02\x82\x01\x01\x00\x8a\x05\xfb\xcb3\xe4E\xe1\xf5\xad\xa2\xcf\x14%\xb5\x92K\x03x\x11\xe1\xe2\xb40\x0fkp\x99{\x10\u023fq\xa2n\f\xd0\x01\x1d\x9a\x98AA\r\x10\xe7CyH\xd7F\xa7\x99\xec\x1bvk\xc0\xfc\xecUlfh\xcd\xden\x98\xdbI\xb4`Ao\xb5\xe9}%\b\xc8K\xc3\xdfX\xeePC=\x17\xf6AD\xa0\r\xd4R\xc85Y\x80I\x82\u0740\xe3:\xf5i\xdb\a'Gu\xea\t\xc9[f\xcf}\x19\xb0\xa5\xab-@\x89\x13P\xb0c\xa8\xbf\xe9X\x87&.\\yi^\xc0ujLA\x85\n\x14\x99\xb3O\xf7\xd0m\x05)\x8aklQ\u0162\xb1AQ\xf3\xa8l\xf9s\xd3}\x1bQ\xdcj\xb0\xd3\xf9\xd9\u00fbK\xf9\x9b\xb6&A\x97\xaf\xe8\xabWt\u007f\xb5W\x84P\u03ce\xd4\xcd;\u07b2\x9f\x16B\xbf4\x06\xe3%\x10\xa7\xce\x1b\xf4\x1e)\xa3=\xcb\x11v\xd7PD\xa2\xf3S\xa9\xbcCG\x0e\r\x940!\x8cZ#\xa6\xe0Sx|\xe2\xcai\xf7\xf1\x02\x81\x81\x00\xec8\xf2\x871\xad\x15\n\xcd;\xfd\xdc\u0599\x1e4\xfax\u02e3S\x80\x14I\xd6|1\xdb \xb3\xb7\xdb\x12\x9a/ej\xacA\x98?8*\xbc\xa1u\xbb\xb1\xbb\u06ac3\xf8X\xe2\x12\x1fhjI\xc3\xd8:\xad]\f\xc54\xab\xa9\xcc&\x1707XV\n\x9e\xf2\u0611\x17~\xe1{\xc8\r\xdb?i\xb9G\x02'\xcbDt\u0096x\x83Qp\x05\xcf+\xbb\u019c\x1b\x84\"}(\x05\xb7\x1b\xec\x13\xe4\xfbC\xe8\x03_\xa83\x02\x81\x81\x00\u019c\x14\xb5x\xa7<\xa5#\xdd\xc7:\x18\xb36\x05\xfe\tbKy\xa3\xae\u0667\u03f3\xad\x1bC\xf3lU\xaf\xbd+Vh\xcd\xf20>\xfe\xfb\x13t\xba\xfe\xc7\xe8\xfc\x84\x13\x18#\x8e\xc0\xd6*`I\x86\x11\u0703\xa5\x96\xf1r\x03\xfa\u06ce\xba\xb9brS\x11\\[\x1c\x9b\x9f\x99O'\xdb'\x00-\xb0\xa4\xf5}#\x8d}\u07f3rU\xf9\x98\u03f2f\x8a\x15\xeb^\x91\x04M\x04\xbeT\xeai\x11\x88\x95\u05d1\xfa!~E\x02\x81\x80\x04\xf4\xc0\xe9\xe39\xedj\x17\x9a=\x9eG\x86X\xe2\xe5\xaai#Y\x1a#\xd2\xd4\xc40K\x97\xa9\r\x9ft\tv\x1b\x1b\x9c\u05d7y$\x15\x89u\x9d\xc1\xbd]\u2760M\x82\x97\xe3\xa7s\n_\xd7\xd3\x0e\x90\x1a\x96\xad\x00\x88\xe5|\x1a\x04|\x87|\xbb\xf9g\x12SF\xe0\x06Rv\xc5`\xbd\v\u070c\xfd\x97f\xfcU\xc8YX\xcc\xd8|y\xec5G\x86\x9e\t\"n\xa8F\x95\xf3`\xbf\x1e9\xe0\xa1\x00\xf0\xba9d\x8f\x02\x81\x81\x00\xae\x85\xf6\xf8\xa4F(\xd7`\x96\x00\xfc\u007f9\xf3\x8d\xfevV\x86\xc6#\u0700\xec\xa2j\\\x02*\x8a\xdc6\u02cf\xeaf+\x00V\x02\xa9H\xbcn\x93S\u0090\xfe\x9d\xee\x9d,\xc8\n(O\u027b\x04\xc0\xa8/c\v~\x81:T^\xfa\xd2\b(\xe1f^\xb2Q\xba\xca\u007f\xc2\x16 \xe5\x80\x01\\y\xc9\u041a\xdb\xd5\x02\xff\x8a4\x90\x93\x16\x1a~\xe0`\x94\x94\xc0X@\fLu\x8f\x9e\xe4\xc1a#\xe9\xb0\x1dYtW\x15\x02\x81\x81\x00\x8b!\xd6\xfdA\xbe\xad\xf6W+6\xe1\x84:XX\xbawCy\xa5u\x02\b\ry\xb6\xe6\xa0\x1eJz3G?&\xa5\x9e2\xb40\xd2\xdf\xe0\xac\a|*\xc3\x15\x8f\xf9\xe0q\x86\xb2U]A#\x1d*]\xe5\u053fG\xe1\xe2\xbc\v\xaa\x8f\xec\xaa!\xf4\x9eT{/$h\x84u\u063d\xc0\u007f\x0f\x06\u01fc>K\x8f3\xa3\xd2\xf3\xac\xfa#\xf0l39\xf1\xf7\x0fr3\xd0\xd0]\x99\x90\xca\xfc\x10\u04ab`}\a&\bF")