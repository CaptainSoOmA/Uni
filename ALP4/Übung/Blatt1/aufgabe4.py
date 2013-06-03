texts = ["Wo ist mein Notebook?","In der Pauseein Apfel.","-nur reine Algorithmik -","Mache gerne Informatik !","Besser ist ein Tutorium.","Alp 4 ist interessant.","das Tutorium am Mittwoch"]

aim = "Maurer ist ein Trottel !"
done = ""

for p in range(0,len(aim)):
	for t in texts:
		if p < len(t):
			if t[p] == aim[p]:
				print p+1,":",t
				break
