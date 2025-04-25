# ProtonMail bridge for git send-email

When submitting a patch to the Linux kernel, or to any other project that uses
mailing lists rather than PRs, using `git send-email` is the best tool for the
job. But if using ProtonMail, you have to get the bridge working first.

## Set up ProtonMail bridge

I had trouble getting KWallet to work nicely with the bridge (the bridge would
say it couldn't access a secret service, and would quit). However, I was able
to get `pass` working. First, generate a GPG key pair. If you already have one,
you can skip this step:

```sh
gpg --full-generate-key
```

To find the key ID:

```sh
gpg -K
```

Install `pass` and `pinentry-tty` (not sure if that one is actually needed, but
it might help if you used a passphrase):

```sh
sudo apt install pass pinentry-tty
```

After installing `pass`, initialize a new store:

```sh
pass init YOURGPGKEYID
```

If you set a passphrase on your key, you need to cache it in the `gpg-agent` so
you aren't prompted. Try decrypting something in your store (you can tab
complete the secret path):

```sh
pass show path/to/secret
```

*Immediately* after running the above, open the ProtonMail bridge and follow
the instructions to authenticate to your account. Once signed in, use the
information on the dashboard to add the following to `~/.gitconfig`:

```sh
# Change the below accoding to the bridge dashboard
[sendemail]
	smtpserver = 127.0.0.1
	smtpuser = email@example.org
	smtpencryption = ''
	smtpserverport = 1025
```

## Send the patch

Commit all of your proposed changes in a single commit, then:

```sh
git format-patch
```

This will create a file in the repo root with the suffix `.patch`. Now, send
the patch:

```sh
git send-email 0001-my-commit.patch
```

The above command will prompt you for the required information, such as who to
send the email to, the password (use the one listed in the dashboard on the
bridge), confirmation, etc. Then it will send the email using your Proton
account via the bridge.

    #mail #email #proton #patch
