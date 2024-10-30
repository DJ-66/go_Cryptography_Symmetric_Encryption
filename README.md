# We're using a 32 byte long seed_String
# seed_String will be used to programmically generate a random 32byte salt string

# Using AES alone, you can only encrypt/decrypt
# data that is 16bytes long or longer(which is the size of an AES block)

# Using AES + GCM (combined mode) aleviates the AES size limitation
# Say if we wanted to use random length user generated strings (often less than 16bytes) 
# i.e.  To store user generated passwords as encrypted string our DB
# Aes + GCM also gives us message authentication (integrity) = Authenticated Encyption

