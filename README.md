# Working AES + GCM encrypt / decrypt for strings of any size (ex. passwords)

# We're using a 32 byte long seed_String
# seed_String is used to programmically generate a random 32byte salt string

# Unique salt string is used by aes.NewCipher to encrypt / decrypt plain text

# Using AES alone, you can only encrypt / decrypt
# data that is 16bytes long or longer (The size of an AES block)

# AES + GCM (combined mode) aleviates the AES size limitation
# Say if we wanted to use random length user generated strings (often less than 16bytes) 
# i.e.  To store user generated password as encrypted string our DB
# Aes + GCM also gives us message authentication (integrity) = Authenticated Encyption

# Work in progress - will add CLI + User [string input] in the future

