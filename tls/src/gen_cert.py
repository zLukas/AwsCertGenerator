from OpenSSL import crypto

def create_self_signed_cert(d: dict, write_files=False):
            
        # create a key pair
        k = crypto.PKey()
        k.generate_key(crypto.TYPE_RSA, 4096)
        # create a self-signed cert
        cert = crypto.X509()
        cert.get_subject().C = d["C"]
        cert.get_subject().ST = d["ST"]
        cert.get_subject().L = d["L"]
        cert.get_subject().O = d["O"]
        cert.get_subject().OU = d["OU"]
        cert.get_subject().CN = d["CN"]
        cert.set_serial_number(d["serial_num"])
        cert.gmtime_adj_notBefore(0)
        cert.gmtime_adj_notAfter(d["expiration_time"])#10*365*24*60*60 s
        cert.set_issuer(cert.get_subject())
        cert.set_pubkey(k)
        cert.sign(k, 'sha1')

        raw_cert = crypto.dump_certificate(crypto.FILETYPE_PEM, cert).decode("utf-8")
        raw_private_key = crypto.dump_privatekey(crypto.FILETYPE_PEM, k).decode("utf-8")
        if write_files:
            open(d["cert_file_name"], "wt").write(raw_cert)
            open(d["key_file_name"], "wt").write(raw_private_key)
            

        return (raw_cert, raw_private_key)

if __name__ == "__main__":
    cert_bundle = create_self_signed_cert()