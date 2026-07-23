*This was a pain in the ass so here is a step-by-step list if you want to renew
your certification after 90 days.*

# How To Renew Certification:

## Renewing Certification

The first step is to disable Nginx, run the following command:

```bash
$ sudo systemctl stop nginx
```

After this, test renewing the certificate using the certbot command below:

```bash
$ sudo certbot renew --dry-run
```

Once this passes, run:

```bash
$ sudo certbot renew
```

After all of this, restart the Nginx server:

```bash
$ sudo systemctl start nginx
```

or 

```bash
$ sudo systemctl restart nginx
```

The certification is now renewed, and you'll be able to access your site. To
test if the certificate worked, run:

```bash
$ openssl s_client -connect terrariadle.net:443 | openssl x509 -noout -dates
```


## Updating Backend Certificates

In the /api directory (backend), there is a /certs folder that contains all the
certifications needed for the backend. You will need to update these.

Run the following set of commands to update the path for the certificates:

```bash
$ sudo cp /etc/letsencrypt/live/terrariadle.net/fullchain.pem /home/ubuntu/Terrariadle/certs/

$ sudo cp /etc/letsencrypt/live/terrariadle.net/privkey.pem /home/ubuntu/Terrariadle/certs/

$ sudo chown ubuntu:ubuntu /home/ubuntu/Terrariadle/certs/*.pem

$ chmod 600 /home/ubuntu/Terrariadle/certs/*.pem
```

This will update the paths in your backend to use the new certificates.

## Closing remarks

If you encounter any more issues, ensure that Nginx is properly formatted.
ChatGPT is a ok resource, but be very specific in the problem and don't
deviate too far from the original source code.

Also, consider making a script that autorenews the certificate, it would save
a lot more headache in the future.