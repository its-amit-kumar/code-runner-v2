a = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
for x in range(10):
    a+=a
f = open("demofile2.txt", "a")
for x in range(10):
    f.write(a)
f.close()
