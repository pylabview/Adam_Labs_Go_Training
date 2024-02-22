package main

var dbData = []Record{
	{
		Title:   "Harry visits King Charles after diagnosis",
		Content: "King seen for first time since cancer announcement, as Duke of Sussex travels from the US.",
		Link:    "https://www.bbc.co.uk/news/uk-68214982?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Boeing: Bolts missing from door, says blowout report",
		Content: "A door that blew away from a Boeing 737 Max may not have been properly secured, a new report says.",
		Link:    "https://www.bbc.co.uk/news/business-68220627?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Michigan school shooter's mother guilty of involuntary manslaughter",
		Content: "Jennifer Crumbley is the first US parent ever convicted of manslaughter over a mass shooting by their child.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68223118?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Romance fraudster conned women out of £300k",
		Content: "He convinced his victims that he was a millionaire but really he was unemployed and living out of a suitcase.",
		Link:    "https://www.bbc.co.uk/news/articles/c04rp592z49o?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Constance Marten and Mark Gordon hid baby's location from police, court hears",
		Content: "In new footage, police officers are heard asking \"where's your child?\", but the couple do not reply.",
		Link:    "https://www.bbc.co.uk/news/uk-68218831?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Dentists to get cash incentives for NHS patients",
		Content: "The move is expected to be part of a wider dental reform plan to boost access for patients in England.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68171168?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Clapham attack: Abdul Shokoor Ezedi is being helped to hide, say police",
		Content: "Police continue to search for corrosive liquid attack suspect Abdul Shokoor Ezedi for a sixth day.",
		Link:    "https://www.bbc.co.uk/news/uk-68216579?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Fake Hull estate agent cannabis growing gang jailed",
		Content: "The five men bought up empty properties in Hull and used a fake estate agent shop as cover.",
		Link:    "https://www.bbc.co.uk/news/uk-england-humber-68222004?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Europe farmers protests: EU scraps plans to halve pesticide use",
		Content: "The announcement appears to be a concession to farmers who have been protesting tighter regulations.",
		Link:    "https://www.bbc.co.uk/news/world-europe-68218907?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Gary Lineker on avoiding 'increasingly toxic' social media platform X",
		Content: "The presenter says changes to the design of X/Twitter means he now sees the \"vitriolic side\" more often.",
		Link:    "https://www.bbc.co.uk/news/entertainment-arts-68215168?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Hertfordshire policeman demoted for selling his trousers",
		Content: "The officer is guilty of gross misconduct and his rank has been reduced following a police hearing.",
		Link:    "https://www.bbc.co.uk/news/uk-england-beds-bucks-herts-68218617?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "The father who never gave up hope on getting justice for his stabbed son",
		Content: "Michael Jonas's dad reflects as the six men who killed his 17-year-old son are handed life sentences.",
		Link:    "https://www.bbc.co.uk/news/articles/cjmp9vylr0jo?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Love Island's Tasha Ghouri wants to normalise deaf accents",
		Content: "The reality TV star says the reaction to a video without her cochlear implant has been \"overwhelming\".",
		Link:    "https://www.bbc.co.uk/news/newsbeat-68085789?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Iain Watson: Labour works on plan for power",
		Content: "Behind the scenes, Labour is working out how it might put its ideas into action if it wins the next general election.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68205102?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Carry On films: The star who helped World War II prisoners escape",
		Content: "Peter Butterworth, who appeared in 16 Carry On films, was an officer for MI9 planning prison escapes.",
		Link:    "https://www.bbc.co.uk/news/entertainment-arts-68209738?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Liz Truss targets 'secret Tories' with new campaign",
		Content: "The ex-PM launches Popular Conservatism to fight for right wing values, without undermining Rishi Sunak.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68218299?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "What's killing so many of Sri Lanka's elephants?",
		Content: "If deaths continue at the current rate, Sri Lanka could lose 70% of its elephants, experts say.",
		Link:    "https://www.bbc.co.uk/news/world-asia-68090450?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Grieving mother exposes the truth of Turkey’s deadly earthquake",
		Content: "After losing her son, Nurgül Göksu set out to investigate how building alterations left his home vulnerable to earthquakes.",
		Link:    "https://www.bbc.co.uk/news/world-europe-68203542?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "A jailed Imran Khan leaves Pakistan divided ahead of election",
		Content: "No other politician has caused as many rifts in Pakistan as the cricket star-turned politician.",
		Link:    "https://www.bbc.co.uk/news/world-asia-68177777?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Super Bowl 58: Taylor Swift conspiracy theories 'nonsense' - NFL boss Roger Goodell",
		Content: "Super Bowl conspiracy theories involving Taylor Swift are \"nonsense\", says NFL commissioner Roger Goodell.",
		Link:    "https://www.bbc.co.uk/sport/american-football/68206676?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Is Ireland's productivity boom real or 'artificial'?",
		Content: "Irelands hosts many multinational companies which makes calculating its true productivity difficult.",
		Link:    "https://www.bbc.co.uk/news/business-67992477?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Black teenager stopped by Met Police 6 times in 5 months",
		Content: "Eight Met officers are under investigation by the police watchdog.",
		Link:    "https://www.bbc.co.uk/news/uk-england-london-68219504?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Shea Gordon: Teens jailed for life after birthday party murder",
		Content: "Shea Gordon, 17, was stabbed just after midnight in 2022 while attending the event in east London.",
		Link:    "https://www.bbc.co.uk/news/uk-england-london-68219942?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "WeWork's ousted boss plots buyback of bankrupt firm",
		Content: "Lawyers for Adam Neumann claim the leaders of his former company are stonewalling his approaches.",
		Link:    "https://www.bbc.co.uk/news/business-68220179?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Tom Holland to star in Romeo and Juliet in West End",
		Content: "The Spider-Man actor last performed in the West End in 2008 as a child in Billy Elliot The Musical.",
		Link:    "https://www.bbc.co.uk/news/entertainment-arts-68219122?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Kwasi Kwarteng to stand down as Spelthorne MP",
		Content: "The former chancellor says it has been an \"honour to serve the residents of Spelthorne since 2010\".",
		Link:    "https://www.bbc.co.uk/news/uk-england-surrey-68214536?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Cash-strapped Birmingham to raise council tax by 10%",
		Content: "The government says it will not block the increase, but that responsibility lies with the council.",
		Link:    "https://www.bbc.co.uk/news/articles/c4nkzejrrx4o?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "PM says pressure starting to ease as millions get last cost-of-living payment",
		Content: "As millions get a final cost of living payment, Rishi Sunak says the burden on households is subsiding.",
		Link:    "https://www.bbc.co.uk/news/business-68204195?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Facebook and Instagram to label all fake AI images",
		Content: "The tech will detect images - but not video or audio - made by other companies' AI generators.",
		Link:    "https://www.bbc.co.uk/news/technology-68215619?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Welsh fire service taken over after harassment probe",
		Content: "The Welsh government effectively takes over as a minister says failures could put lives at risk.",
		Link:    "https://www.bbc.co.uk/news/uk-wales-politics-68217579?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Paul Mackenzie: Kenyan cult leader charged with 191 murders",
		Content: "Paul Mackenzie, accused of starving his followers to death, pleads not guilty.",
		Link:    "https://www.bbc.co.uk/news/world-africa-68214749?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Churchill's false teeth 'snapped up' for £18,000",
		Content: "A gold-mounted set of false teeth made for the prime minister sell at auction in Cheltenham for £18,000.",
		Link:    "https://www.bbc.co.uk/news/uk-england-gloucestershire-68215240?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Houthis claim new attacks on Red Sea shipping",
		Content: "The latest attacks come in defiance of repeated US and UK strikes to try to deter the group.",
		Link:    "https://www.bbc.co.uk/news/world-middle-east-68218901?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Christian Horner allegations: Hearing into complaint against Red Bull team principal set for Friday",
		Content: "A hearing into a complaint of inappropriate behaviour made against Red Bull Formula 1 team boss Christian Horner will take place on Friday.",
		Link:    "https://www.bbc.co.uk/sport/formula1/68218956?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Spanish farmers join wave of protests",
		Content: "Across Spain, farmers have been protesting at EU regulations and demanding more help from their government.",
		Link:    "https://www.bbc.co.uk/news/world-europe-68216353?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Tesla owners told not to wear Apple virtual reality headsets while driving",
		Content: "Videos have emerged showing people wearing the virtual reality headsets while in self-driving cars.",
		Link:    "https://www.bbc.co.uk/news/technology-68215614?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Israel-Gaza war: Unknown fate of six-year-old Hind Rajab trapped under fire",
		Content: "Six-year-old Hind was last heard from trapped in a car surrounded by bodies - then the line went dead.",
		Link:    "https://www.bbc.co.uk/news/world-middle-east-68180642?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "RAF servicewoman hid in toilet to escape sexual harassment from boss",
		Content: "Servicewomen tell the BBC they quit the air force after being \"seen as trouble\" for complaining.",
		Link:    "https://www.bbc.co.uk/news/uk-england-lincolnshire-68168403?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Karolina Shiino: Ukraine-born Miss Japan gives up crown following affair",
		Content: "The controversial Ukraine-born model has resigned after a report on her relations with a married man.",
		Link:    "https://www.bbc.co.uk/news/world-asia-68213223?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Los Angeles: Helicopter crew rescues man from surging river",
		Content: "The man entered the water to save his dog after it was swept away. The dog managed to swim to safety.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68214577?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Former Irish PM John Bruton dies aged 76",
		Content: "Mr Bruton was taoiseach from 1994 to 1997 when he led a rainbow coalition government.",
		Link:    "https://www.bbc.co.uk/news/articles/cd18yqnke9no?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Online Safety Act not strong enough, Brianna Ghey's mum says",
		Content: "Esther Ghey says the act does not go far enough to protect children and calls for \"drastic action\".",
		Link:    "https://www.bbc.co.uk/news/uk-england-manchester-68214691?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "NatWest may 'tell Sid' to buy shares in June",
		Content: "The firm overseeing the government's stake in the bank may soon begin selling shares to the public.",
		Link:    "https://www.bbc.co.uk/news/business-68216611?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Ukraine war: Baby killed in Russian strike on Kharkiv hotel",
		Content: "A two-month-old boy is killed in the latest cross-border attack on Ukraine's Kharkiv region.",
		Link:    "https://www.bbc.co.uk/news/world-europe-68214631?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Jaywick dog attack: Esther Martin will not be forgotten, says son",
		Content: "Esther Martin, 68, died after being attacked in a home on the Essex coast.",
		Link:    "https://www.bbc.co.uk/news/uk-england-essex-68216175?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Taylor Swift reveals new album details as Super Bowl speculation continues",
		Content: "The popstar announced the release of her 11th album in a Grammy Awards acceptance speech on Sunday.",
		Link:    "https://www.bbc.co.uk/news/entertainment-arts-68207651?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Pterosaur: Unique flying reptile soared above Isle of Skye",
		Content: "The pterosaur soared over the heads of dinosaurs in Scotland 168-166 million years ago.",
		Link:    "https://www.bbc.co.uk/news/science-environment-68207021?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Toby Keith: Country music legend dies aged 62",
		Content: "Toby Keith, one of country music's biggest stars, had previously revealed he had stomach cancer.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68215639?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Trent Alexander-Arnold: Liverpool defender discusses life outside football with academy players",
		Content: "Liverpool defender Trent Alexander-Arnold speaks to Liverpool academy players about the mental pressures that come with being released by clubs.",
		Link:    "https://www.bbc.co.uk/sport/av/football/68209616?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Rishi Sunak says he was taken by surprise on £1,000 Rwanda bet",
		Content: "Rishi Sunak says he took the bet with Piers Morgan to \"underline\" his commitment to the Rwanda policy.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68214641?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Sheffield: New mum urges HIV testing amid case rise",
		Content: "Becky has spoken out as official figures show testing for HIV remains below pre-pandemic levels.",
		Link:    "https://www.bbc.co.uk/news/uk-england-south-yorkshire-68144972?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Yandex: Owner of 'Russia's Google' pulls out of home country",
		Content: "Yandex has previously been accused of hiding information from Russia's public about the Ukraine war.",
		Link:    "https://www.bbc.co.uk/news/business-68213191?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "NHS Wales: Day-long ambulance wait after care home fall",
		Content: "The family of Theresa Jones, 91, say she was left on the floor \"like a piece of rubbish\".",
		Link:    "https://www.bbc.co.uk/news/uk-wales-68211110?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Scottish comedy 'titans' to star in new BBC Scotland sitcom",
		Content: "Greg McHugh and Gregor Fisher will play a father and son duo in the BBC Scotland comedy Only Child.",
		Link:    "https://www.bbc.co.uk/news/uk-scotland-68209924?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "McDonald's sales dented by Israel-Gaza boycotts",
		Content: "The chain said its fourth quarter sales failed to meet its target for the first time in four years.",
		Link:    "https://www.bbc.co.uk/news/business-68209085?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Clapham attack: Man arrested for allegedly helping Abdul Shokoor Ezedi",
		Content: "Police say a man was detained and bailed on suspicion of helping Abdul Shokoor Ezedi.",
		Link:    "https://www.bbc.co.uk/news/uk-68202153?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Jaywick dog attack: Esther Martin's death causes 'terrible shock'",
		Content: "Jaywick resident Mary Donovan says it is a \"terrible shock\" for everyone in the coastal village.",
		Link:    "https://www.bbc.co.uk/news/uk-england-essex-68209217?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Usyk missed birth of child training for postponed Fury fight",
		Content: "Oleksandr Usyk reveals he missed the birth of his second child while training in Spain for the fight with Tyson Fury that has been postponed.",
		Link:    "https://www.bbc.co.uk/sport/boxing/68216146?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Andy Murray's disappointing season continues with loss at Open 13 Provence in Marseille",
		Content: "Andy Murray's disappointing start to the year continues as he is beaten by Czech world number 66 Tomas Machac at the Open 13 Provence.",
		Link:    "https://www.bbc.co.uk/sport/tennis/68219931?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Jordan 2-0 South Korea: Mousa Tamari scores to send Jordan into first Asian Cup final",
		Content: "Mousa Tamari scores a superb solo goal to help stun South Korea and send Jordan to their first Asian Cup final.",
		Link:    "https://www.bbc.co.uk/sport/football/68219853?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "'I will not leave you alone' - Oleksandr Usyk sends message to Tyson Fury",
		Content: "Oleksandr Usyk opens up to BBC Sport about his immediate reaction to the news his heavyweight unification fight with Tyson Fury was postponed, missing the birth of his daughter and his message to the British fighter.",
		Link:    "https://www.bbc.co.uk/sport/av/boxing/68220168?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "King Charles's cancer was 'caught early', Rishi Sunak says",
		Content: "Prime minister tells the BBC he is in \"regular contact\" with King Charles III following his diagnosis.",
		Link:    "https://www.bbc.co.uk/news/uk-68213383?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Prince William can show his own version of royalty",
		Content: "The King's illness will make the Prince of Wales's role as heir even more important.",
		Link:    "https://www.bbc.co.uk/news/uk-68219415?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "William, Harry and the other royals - what it means for them",
		Content: "It's been a bleak midwinter for the Royal Family. Will the King's health news help to bring them together?",
		Link:    "https://www.bbc.co.uk/news/uk-68211941?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Why did Harry and Meghan leave the Royal Family and where do they get their money?",
		Content: "The Duke and Duchess of Sussex are not funded by the Royal Family and have various commercial deals.",
		Link:    "https://www.bbc.co.uk/news/explainers-51047186?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "What do we know about the diagnosis?",
		Content: "King Charles has been diagnosed with cancer, but Buckingham Palace has not said what type it is.",
		Link:    "https://www.bbc.co.uk/news/health-68203457?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "How the King's duties will change",
		Content: "The King is stepping back from public duties after being diagnosed with cancer.",
		Link:    "https://www.bbc.co.uk/news/uk-56201331?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "The work that will continue during the King's treatment",
		Content: "Charles is stepping back from public-facing duties, but his private work continues, BBC political editor Chris Mason writes.",
		Link:    "https://www.bbc.co.uk/news/uk-68213215?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Prince Harry to fly to UK to visit his father",
		Content: "The Duke of Sussex spoke to his father about the diagnosis and will travel to the UK to visit him.",
		Link:    "https://www.bbc.co.uk/news/uk-68209692?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Buckingham Palace statement in full",
		Content: "Buckingham Palace has announced that King Charles is being treated for a form of cancer.",
		Link:    "https://www.bbc.co.uk/news/uk-68205659?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Watch BBC News read Palace statement as King's cancer diagnosis announced",
		Content: "Buckingham Palace has announced that the King is being treated for cancer.",
		Link:    "https://www.bbc.co.uk/news/uk-68209896?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Cost of living payment: When is it and who gets it?",
		Content: "Millions of low-income households will receive a final cost of living payment soon.",
		Link:    "https://www.bbc.co.uk/news/business-61592496?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Fuel cost: How to save petrol and diesel",
		Content: "With weak competition adding to fuel prices, are drivers doing the right things to save money?",
		Link:    "https://www.bbc.co.uk/news/business-61745697?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Train strikes: All you need to know about action in February and March",
		Content: "More than 300 London Overground workers will go on strike from midnight on Monday 19 February.",
		Link:    "https://www.bbc.co.uk/news/business-61634959?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Energy bills: What can I do if I can't afford to pay?",
		Content: "If you're struggling to afford your gas and electricity bills, what options are available?",
		Link:    "https://www.bbc.co.uk/news/business-62435432?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "National Insurance calculator: What will I pay and how is tax changing?",
		Content: "National Insurance rates are falling, but other changes mean many people are paying more tax.",
		Link:    "https://www.bbc.co.uk/news/explainers-63635185?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "How to get a job: Six expert tips for finding work",
		Content: "There are 10m people out of work in the UK, so if you're searching for a job you're not alone. Here are some tips on how to get started.",
		Link:    "https://www.bbc.co.uk/news/business-64939070?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Mortgage rates: Five ways to save money",
		Content: "Experts give advice for those who might be worried about their monthly mortgage payments.",
		Link:    "https://www.bbc.co.uk/news/business-65984415?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Rent increases: What you can do about a rise from your landlord",
		Content: "The BBC's Lora Jones tells you four things you can do, if your landlord asks for more money.",
		Link:    "https://www.bbc.co.uk/news/business-63290805?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Renting: What are your rights as a tenant?",
		Content: "With one in five people now renting in the UK, it's important to understand your rights as a tenant.",
		Link:    "https://www.bbc.co.uk/news/technology-65038459?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "How to save money even when the budget is tight",
		Content: "The interest paid on savings is better than anything seen for years, so how can you save when bills are rising?",
		Link:    "https://www.bbc.co.uk/news/business-66855566?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Ukraine war: Baby killed in Russian strike on Kharkiv hotel",
		Content: "A two-month-old boy is killed in the latest cross-border attack on Ukraine's Kharkiv region.",
		Link:    "https://www.bbc.co.uk/news/world-europe-68214631?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Michigan school shooter's mother guilty of involuntary manslaughter",
		Content: "Jennifer Crumbley is the first US parent ever convicted of manslaughter over a mass shooting by their child.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68223118?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Boeing safety checks to use 'more boots on ground'",
		Content: "The head of the US aviation regulator told US lawmakers they were stepping up safety checks.",
		Link:    "https://www.bbc.co.uk/news/business-68220625?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Los Angeles: Helicopter crew rescues man from surging river",
		Content: "The man entered the water to save his dog after it was swept away. The dog managed to swim to safety.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68214577?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Rio de Janeiro: Dengue spike prompts health emergency ahead of Carnival",
		Content: "The city is seeing a huge rise in the cases of dengue fever just days before its colourful Carnival.",
		Link:    "https://www.bbc.co.uk/news/world-latin-america-68215360?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Tucker Carlson to interview Russia's Putin",
		Content: "The ex-Fox News host will be the first Western journalist to interview Vladimir Putin since Russia invaded Ukraine.",
		Link:    "https://www.bbc.co.uk/news/world-europe-68223148?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Karolina Shiino: Ukraine-born Miss Japan gives up crown following affair",
		Content: "The controversial Ukraine-born model has resigned after a report on her relations with a married man.",
		Link:    "https://www.bbc.co.uk/news/world-asia-68213223?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Donald Trump does not have presidential immunity, US court rules",
		Content: "The former president loses a landmark legal bid to shield him from charges of election fraud.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68026175?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Chile wildfires: Hundreds missing as thousands of homes burn",
		Content: "More than 120 people have been killed in the fires and the number of dead is expected to rise further.",
		Link:    "https://www.bbc.co.uk/news/world-latin-america-68215354?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Spanish farmers join wave of protests",
		Content: "Across Spain, farmers have been protesting at EU regulations and demanding more help from their government.",
		Link:    "https://www.bbc.co.uk/news/world-europe-68216353?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Houthis claim new attacks on Red Sea shipping",
		Content: "The latest attacks come in defiance of repeated US and UK strikes to try to deter the group.",
		Link:    "https://www.bbc.co.uk/news/world-middle-east-68218901?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Paul Mackenzie: Kenyan cult leader charged with 191 murders",
		Content: "Paul Mackenzie, accused of starving his followers to death, pleads not guilty.",
		Link:    "https://www.bbc.co.uk/news/world-africa-68214749?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Super Bowl 58: Taylor Swift conspiracy theories 'nonsense' - NFL boss Roger Goodell",
		Content: "Super Bowl conspiracy theories involving Taylor Swift are \"nonsense\", says NFL commissioner Roger Goodell.",
		Link:    "https://www.bbc.co.uk/sport/american-football/68206676?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "John Cage: Organ playing 639-year-long piece changes chord",
		Content: "Composed by avant-garde artist John Cage, the piece is expected to play in Germany until the year 2640.",
		Link:    "https://www.bbc.co.uk/news/world-europe-68209691?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Yandex: Owner of 'Russia's Google' pulls out of home country",
		Content: "Yandex has previously been accused of hiding information from Russia's public about the Ukraine war.",
		Link:    "https://www.bbc.co.uk/news/business-68213191?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Police to pay $1.9m to black family held at gunpoint in Colorado",
		Content: "The 2020 stop in Colorado was recorded by bystanders, which showed children handcuffed and crying.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68211285?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "El Salvador election: Nayib Bukele revels in landslide win",
		Content: "Preliminary results suggest the incumbent has won more than 80% of the vote.",
		Link:    "https://www.bbc.co.uk/news/world-latin-america-68205036?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Toby Keith: Country music legend dies aged 62",
		Content: "Toby Keith, one of country music's biggest stars, had previously revealed he had stomach cancer.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68215639?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Footballer Dani Alves's rape trial begins in Spain",
		Content: "The ex-Brazil and Barcelona defender denies the assault, saying the sex was consensual.",
		Link:    "https://www.bbc.co.uk/news/world-europe-68208729?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Tesla owners told not to wear Apple virtual reality headsets while driving",
		Content: "Videos have emerged showing people wearing the virtual reality headsets while in self-driving cars.",
		Link:    "https://www.bbc.co.uk/news/technology-68215614?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Nikki Haley seeks Secret Service protection, citing 'multiple issues'",
		Content: "The last challenger to Donald Trump says her campaign has had \"multiple issues\" with security.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68213181?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Theo O Ebonyi: Nigerian pastor arrested over allegedly swindling followers of $1m",
		Content: "Theo O Ebonyi, accused of swindling his followers and others, dismisses news of his arrest as fake.",
		Link:    "https://www.bbc.co.uk/news/world-africa-68163590?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Atmospheric river: Three killed as record rainfall drenches California",
		Content: "The storm brings landslides and mudslides to Los Angeles and San Francisco and an avalanche near Las Vegas.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68202944?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Senegal on the brink after elections postponed",
		Content: "Senegal is seen as a bastion of democracy in West Africa, yet its polls were delayed with three weeks to go.",
		Link:    "https://www.bbc.co.uk/news/world-africa-68209178?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "McDonald's sales dented by Israel-Gaza boycotts",
		Content: "The chain said its fourth quarter sales failed to meet its target for the first time in four years.",
		Link:    "https://www.bbc.co.uk/news/business-68209085?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Police apologise to alleged victim in ice hockey sexual assault case",
		Content: "It took years for Canadian police to lay charges against five current or ex-NHL players accused of sexual assault.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68211284?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Injured, hungry and alone - the Gazan children orphaned by war",
		Content: "Some 19,000 children have been orphaned by the war, or have no adults to look after them, says the UN.",
		Link:    "https://www.bbc.co.uk/news/world-middle-east-68141039?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "At least half of Gaza's buildings damaged or destroyed, new analysis shows",
		Content: "Satellite data reveals at least 144,000 hit since Israeli strikes began, and how south is now bearing the brunt.",
		Link:    "https://www.bbc.co.uk/news/world-middle-east-68006607?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Hamas says it is studying proposal for new pause in Gaza fighting",
		Content: "A framework reportedly sets out a six-week pause in the fighting to allow more hostages to be freed.",
		Link:    "https://www.bbc.co.uk/news/world-middle-east-68138853?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Israeli forces kill three Palestinian fighters in West Bank hospital raid",
		Content: "CCTV captures undercover forces disguised as medics and nurses in the hospital with weapons raised.",
		Link:    "https://www.bbc.co.uk/news/world-middle-east-68137050?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "UNRWA: Gaza aid agency says it is 'extremely desperate' after funding halted",
		Content: "More countries halt aid to UNRWA over the alleged role of some staff in the 7 October Hamas attacks.",
		Link:    "https://www.bbc.co.uk/news/world-middle-east-68127544?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Israel-Gaza war: Counting the destruction of religious sites",
		Content: "Some of the world’s oldest churches and mosques have been completely destroyed in the conflict.",
		Link:    "https://www.bbc.co.uk/news/world-middle-east-67983018?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Orcas gasp for air whilst trapped in drift ice off the coast of Japan",
		Content: "The coast guard is unable to rescue them as the ice is too thick, an official told local media.",
		Link:    "https://www.bbc.co.uk/news/world-68222247?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Drone video shows graffiti covering luxury high-rise in LA",
		Content: "Police say 27 floors were painted after over a dozen people entered the abandoned complex last week.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68209951?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Drone footage shows Chile forest fire devastation",
		Content: "At least 112 people have been killed in Chile's Valparaiso region.",
		Link:    "https://www.bbc.co.uk/news/world-latin-america-68205132?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Chile wildfires: 99 dead as wildfire tears through Valparaiso region",
		Content: "Officials have warned the death toll is likely to rise as rescue teams reach the hardest hit areas.",
		Link:    "https://www.bbc.co.uk/news/world-latin-america-68197170?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "US strikes on Iraqi town of al-Qaim filmed",
		Content: "Video verified by the BBC shows US strikes on al-Qaim, Iraq.",
		Link:    "https://www.bbc.co.uk/news/world-middle-east-68194714?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Ros Atkins on… Iran, its proxies and the ‘Axis of Resistance’",
		Content: "The BBC's analysis editor explains the Iran-backed militia groups at the centre of US strikes.",
		Link:    "https://www.bbc.co.uk/news/world-68176257?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Michael Jordan championship trainers sell for $8m",
		Content: "The auction of Jordan's shoes from six NBA finals has set a new world record for game-worn trainers.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68172218?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Turkey earthquake: How a grieving mother uncovered the truth behind her son's death",
		Content: "After losing her son, Nurgül Göksu set out to investigate how building alterations left his home vulnerable to earthquakes.",
		Link:    "https://www.bbc.co.uk/news/world-europe-68203542?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Israel-Gaza war: Unknown fate of six-year-old Hind Rajab trapped under fire",
		Content: "Six-year-old Hind was last heard from trapped in a car surrounded by bodies - then the line went dead.",
		Link:    "https://www.bbc.co.uk/news/world-middle-east-68180642?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Nevada caucuses or primary: Why both Trump and Haley may claim victory",
		Content: "Republicans are holding two separate ballots this week, with the candidates not going head to head.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68208325?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Sri Lanka: What's killing so many of the country's iconic elephants?",
		Content: "If deaths continue at the current rate, Sri Lanka could lose 70% of its elephants, experts say.",
		Link:    "https://www.bbc.co.uk/news/world-asia-68090450?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "A jailed Imran Khan leaves Pakistan divided ahead of election",
		Content: "No other politician has caused as many rifts in Pakistan as the cricket star-turned politician.",
		Link:    "https://www.bbc.co.uk/news/world-asia-68177777?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Is Ireland's productivity boom real or 'artificial'?",
		Content: "Irelands hosts many multinational companies which makes calculating its true productivity difficult.",
		Link:    "https://www.bbc.co.uk/news/business-67992477?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "The highs and lows of Grammys 2024 - and why Taylor Swift won album of year",
		Content: "The best (and worst) performances, the memorable moments, and why Swift's big win was inevitable.",
		Link:    "https://www.bbc.co.uk/news/entertainment-arts-68200931?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Huge atom-smasher bid to find missing 95% of Universe",
		Content: "Researchers want a new, much bigger supercollider but is it worth us paying the £12bn price tag?",
		Link:    "https://www.bbc.co.uk/news/science-environment-68172162?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "US Senate releases deal on border and Ukraine - but will it ever become law?",
		Content: "The bill significantly overhauls border and asylum processes, but House Republicans have rejected it.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68143887?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Universal Studios UK: Is a theme park coming to Stewartby?",
		Content: "What might a major new theme park mean for a Bedfordshire village and its 1,200 inhabitants?",
		Link:    "https://www.bbc.co.uk/news/uk-england-68166514?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Family life: Mum rejects social media perfection for mayhem",
		Content: "Mum-of-five Jessica Hymus-Gant shares the good as well as the \"grotty\" reality of family life.",
		Link:    "https://www.bbc.co.uk/news/uk-wales-68139914?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Dentists to get cash incentives for NHS patients",
		Content: "The move is expected to be part of a wider dental reform plan to boost access for patients in England.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68171168?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Liz Truss targets 'secret Tories' with new campaign",
		Content: "The ex-PM launches Popular Conservatism to fight for right wing values, without undermining Rishi Sunak.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68218299?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Rishi Sunak says he was taken by surprise on £1,000 Rwanda bet",
		Content: "Rishi Sunak says he took the bet with Piers Morgan to \"underline\" his commitment to the Rwanda policy.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68214641?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Kwasi Kwarteng to stand down as Spelthorne MP",
		Content: "The former chancellor says it has been an \"honour to serve the residents of Spelthorne since 2010\".",
		Link:    "https://www.bbc.co.uk/news/uk-england-surrey-68214536?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Birmingham's 10% tax rise approved amid financial crisis",
		Content: "The government says it will not block the increase, but that responsibility lies with the council.",
		Link:    "https://www.bbc.co.uk/news/articles/c4nkzejrrx4o?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Labour works on plan for power",
		Content: "Behind the scenes, Labour is working out how it might put its ideas into action if it wins the next general election.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68205102?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Electric cars: Lords urge action on 'misinformation' in press",
		Content: "The Lords Climate Change Committee called on the government to push back against mistruths on range and cost.",
		Link:    "https://www.bbc.co.uk/news/science-environment-68130432?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Prime Minister Rishi Sunak claims cost of living pressures starting to ease",
		Content: "As millions get a final cost of living payment, Rishi Sunak says the burden on households is subsiding.",
		Link:    "https://www.bbc.co.uk/news/business-68204195?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Fines for missing heat pump targets could be dropped, Downing Street hints",
		Content: "Plans to penalise boiler makers for missing heat pump installation targets are reportedly being scrapped.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68207355?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Labour pledges full equal pay rights for ethnic minorities",
		Content: "The party says it would strengthen protections for ethnic minorities and disabled people if it wins power.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68202541?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Martyn's Law: Consultation into anti-terror legislation launched",
		Content: "The proposals would require venues to take \"necessary but proportionate steps\" to protect the public.",
		Link:    "https://www.bbc.co.uk/news/uk-england-manchester-68208227?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Rishi Sunak admits he has failed to cut NHS waiting lists",
		Content: "The PM is also challenged to make a bet on another of his key pledges, of sending some asylum seekers to Rwanda.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68202148?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Assistance dog refusal to be tackled by government",
		Content: "The long-awaited Disability Action Plan also includes support for \"aspiring disabled politicians\".",
		Link:    "https://www.bbc.co.uk/news/articles/cv2jq3m42geo?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Former Scottish Labour leader Kezia Dugdale admits voting SNP",
		Content: "Kezia Dugdale tells the BBC she backed the party at the 2019 European Parliament election because of Brexit.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68185435?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Abuse of councillors and staff putting democracy at risk, say local government groups",
		Content: "Stalking, threats and assaults are putting people off standing for election, local government groups warn.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68167190?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Gillian Keegan: I'm confident of 15 hours of free childcare by April",
		Content: "But the education secretary does not confirm the scheme will be ready for all eligible parents.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68197595?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Grant Shapps says UK and US strikes on Yemen Houthis 'not an escalation'",
		Content: "More than 30 targets were hit in Yemen in the third round of joint strikes on the Iran-backed group.",
		Link:    "https://www.bbc.co.uk/news/uk-68196192?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Sir Bob Neill: Former Tory minister to stand down at next general election",
		Content: "Sir Bob Neill has said he wants to spend more time with his wife, who suffered a stroke in 2019.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68190390?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "UK armed forces not ready for high-intensity war, MPs warn",
		Content: "The armed forces are losing people faster than they can recruit them, says committee of MPs.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68181275?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Stormont: Michelle O'Neill makes history as nationalist first minister",
		Content: "The Sinn Féin vice-president hails \"a new dawn\" as a power-sharing government returns in NI.",
		Link:    "https://www.bbc.co.uk/news/uk-northern-ireland-politics-68180505?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Britain ‘full of secret Conservatives’ - Liz Truss",
		Content: "Launching the Popular Conservative (PopCon) group, the former PM says that \"democracy has become unfashionable\".",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68219261?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Houthis believe they are the region’s Robin Hood - Shapps",
		Content: "The UK defence secretary says Houthi actions are damaging international economies, affecting food prices and have “rocketed” insurance costs.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68210488?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Badenoch claims media is 'trying to fuel gossip'",
		Content: "The business secretary blames the media for reports of a plot for her to replace Rishi Sunak as a “nonsense story”.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68170890?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "MP: 'Garrote' offenders with intestines to cut fly tipping",
		Content: "Desmond Swayne suggests drastic action to cut the number of fly tipping cases.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68167504?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Rachel Reeves quizzed on bankers bonus U-turn",
		Content: "The BBC's business editor asks how Labour can be trusted after its drops its opposition to lifting the cap on bankers’ bonuses.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68167505?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Chris Mason: King to continue some work during cancer treatment",
		Content: "Charles is stepping back from public-facing duties, but his private work continues, BBC political editor Chris Mason writes.",
		Link:    "https://www.bbc.co.uk/news/uk-68213215?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Can 'super libraries' survive spending cuts?",
		Content: "Libraries are transforming what they offer, but will that be enough to save them?",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68168034?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Has Labour dropped its £28bn green investment plan?",
		Content: "The party's leadership is sending mixed messages on what was once their flagship election policy.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68170566?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Laura Kuenssberg: When will politicians make their minds up on key plans?",
		Content: "No wonder voters are unconvinced when politicians may seem unclear on economic plans, writes Laura Kuenssberg.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68192206?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "General election: Why the battle for votes in Scotland is key",
		Content: "Labour aims for an election comeback from one seat in Scotland, the BBC's Nick Eardley writes.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68166148?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "General election 2024 poll tracker: How do the parties compare?",
		Content: "How do people say they will vote in the UK general election? Our poll tracker measures the trends.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68079726?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "NI's government has returned Stormont - what you need to know",
		Content: "The timeline on how an executive is being formed with a Sinn Féin first minister for the first time.",
		Link:    "https://www.bbc.co.uk/news/uk-northern-ireland-67726389?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Wellingborough by-election: Could Reform or Greens make a splash?",
		Content: "Labour, Liberal Democrats, Tories are also on ballot papers for the 15 February by-election.",
		Link:    "https://www.bbc.co.uk/news/uk-england-northamptonshire-68076729?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Brexit: What is the new Northern Ireland trade deal?",
		Content: "Ministers have set out new trade arrangements for Northern Ireland, the rest of the UK and the EU.",
		Link:    "https://www.bbc.co.uk/news/explainers-53724381?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "The tough choices facing cash-strapped councils",
		Content: "From switching off CCTV to closing sports centres, councils are weighing up where they can make savings.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68071570?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Trade deals: What has the UK done since Brexit?",
		Content: "The UK formally suspends trade talks with Canada, the first time it has done so with a potential partner since Brexit.",
		Link:    "https://www.bbc.co.uk/news/uk-47213842?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "UK citizen army: Preparing the 'pre-war generation' for conflict",
		Content: "A senior general stops short of backing conscription - but wants the UK to prepare for wider European war.",
		Link:    "https://www.bbc.co.uk/news/uk-68097048?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Who are the Houthis and why are they attacking Red Sea ships?",
		Content: "The Houthis in Yemen have vowed to attack all ships in the Red Sea that are linked to Israel.",
		Link:    "https://www.bbc.co.uk/news/world-middle-east-67614911?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Kingswood by-election: Who can you vote for?",
		Content: "Voting will take place in the by-election on 15 February after MP Chris Skidmore quit.",
		Link:    "https://www.bbc.co.uk/news/articles/cv2l22j0398o?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "The candidates set for Wellingborough by-election",
		Content: "A recall petition led to the ousting of Peter Bone as an MP after he was suspended from the Commons.",
		Link:    "https://www.bbc.co.uk/news/uk-england-northamptonshire-67763457?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "North Sea oil and gas claims fact-checked",
		Content: "The government wants to guarantee annual oil and gas licensing rounds to \"improve energy security\".",
		Link:    "https://www.bbc.co.uk/news/science-environment-67945281?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Laura Kuenssberg: Why is Humza Yousaf predicting 'inevitable' Keir Starmer victory?",
		Content: "Polls suggest Humza Yousaf is yet to convince the Scottish public he's an effective leader - could that change?",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68041203?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Can Rishi Sunak halt growing Tory poll panic?",
		Content: "The PM got a win when his Rwanda bill cleared the Commons, but disquiet is still growing among his MPs.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68029051?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "West Midlands mayor: Who is running, what is the role?",
		Content: "Everything you need to know about the 2024 West Midlands mayoral election.",
		Link:    "https://www.bbc.co.uk/news/articles/cpvrdpv2ekeo?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Is Rishi Sunak increasing taxes on all flights?",
		Content: "Labour says on X that tax is going up for all flights except for private jets, but that's wrong.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-68032617?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "General election: Labour would need record swing to win",
		Content: "Analysis shows the effect of the new election boundaries. Use our tool to look up your constituency.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-67361138?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Reluctant statesman Rishi Sunak turns focus to foreign affairs",
		Content: "The PM has become an engaged international actor, showing his hawkish side in Ukraine and the Red Sea.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-67884751?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Labour Party 'campaign bible' gives hints of general election strategy",
		Content: "The party issues instructions to candidates on how to get their core policies across to voters.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-67993311?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Welsh independence: Leaving UK is viable, says new report",
		Content: "But experts warns Wales would face a \"significant\" challenge raising enough tax revenue.",
		Link:    "https://www.bbc.co.uk/news/uk-wales-politics-67949443?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Rishi Sunak's five promises: What progress has he made?",
		Content: "The prime minister said five priorities should be used to hold his government to account.",
		Link:    "https://www.bbc.co.uk/news/65647308?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Lee Anderson: Outspoken MP who quit as deputy chair",
		Content: "The MP for Ashfield is known for his controversial views but supporters believe he connects with voters.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-64582994?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "General election: When is the next one and who decides?",
		Content: "A UK general election has to be held by 28 January 2025, but could take place earlier.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-62064552?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Rwanda: The main Conservative Party factions jostling for influence",
		Content: "The \"five families\" and other groups of Tory MPs preparing to vote on Rishi Sunak's Rwanda plan.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-67658700?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Kuenssberg: The thorny politics of Houthi strikes for Sunak and Starmer",
		Content: "The PM and the Labour leader agree on military action but it could be just the beginning of their problems.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-67843756?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "The key NHS targets that have never been met",
		Content: "BBC News finds key NHS targets have been missed for at least seven years - and two have never been met.",
		Link:    "https://www.bbc.co.uk/news/health-67884322?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "National Insurance: Is Jeremy Hunt right about £1,000 saving?",
		Content: "Chancellor Jeremy Hunt claims National Insurance cut is worth 'nearly £1,000 for a typical two-earner family'.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-67912689?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Chris Mason: Justice at last, but plenty more questions remain",
		Content: "Figures across politics will be asked why exoneration for postmasters took so long, our political editor writes.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-67942277?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "National Insurance calculator: What will I pay and how is tax changing?",
		Content: "National Insurance rates are falling, but other changes mean many people are paying more tax.",
		Link:    "https://www.bbc.co.uk/news/explainers-63635185?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Prime Minister's Questions: Which subjects dominated in 2023?",
		Content: "Ahead of the first Prime Minister's Questions of 2024, we look back at what came up most last year.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-67893447?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Would Labour spend £28bn a year on green projects?",
		Content: "What is Labour's green investment plan and what would it do to the UK economy?",
		Link:    "https://www.bbc.co.uk/news/67801144?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Chris Mason: Has Sunak really ruled out a spring election?",
		Content: "The PM's curious choice of words suggests he may still be keeping his options open.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-67883767?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Scottish politics gears up for general election",
		Content: "Parties are making pitches to voters after the prime minister signalled when he will call an election.",
		Link:    "https://www.bbc.co.uk/news/uk-scotland-scotland-politics-67886864?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Was minister right in veterans row with Vorderman?",
		Content: "Carol Vorderman says Johnny Mercer has not met his pledge to veterans - the minister says he has.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-67842826?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Government's asylum figures show uncleared backlog",
		Content: "Thousands of asylum applications made before June 2022 have still not received initial decisions.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-67863380?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Louise Casey: Lessons for government in election year",
		Content: "Trouble-shooter and peer Louise Casey looks at the way to get things done in government for a new BBC series.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-67833145?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Fears UK not ready for deepfake general election",
		Content: "A senior Tory MP leads calls for more government action to prevent AI sabotaging British democracy.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-67518511?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Laura Kuenssberg: Five facts from a political year of gains and losses",
		Content: "Departures and by-elections, polls and Partygate fallout - how the political year unfolded.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-67737116?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Penny Mordaunt says Boris Johnson's Covid WhatsApp messages went missing",
		Content: "MP Penny Mordaunt tells the Covid inquiry messages to Boris Johnson and Michael Gove went missing.",
		Link:    "https://www.bbc.co.uk/news/technology-67780595?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Drakeford resignation: Who will be Wales' next first minister?",
		Content: "BBC Wales political editor Gareth Lewis looks at names in the mix to be next Welsh Labour leader.",
		Link:    "https://www.bbc.co.uk/news/uk-wales-politics-67692529?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "New-year honours: How does the UK honours system work?",
		Content: "How are the new-year, King's birthday and prime minister's resignation honours decided and awarded?",
		Link:    "https://www.bbc.co.uk/news/uk-11990088?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "What is a recession and how could it affect me?",
		Content: "New figures raise the possibility that the UK's economy could fall into recession.",
		Link:    "https://www.bbc.co.uk/news/business-52986863?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Addicted to division: Sunak beset by warring Tory factions",
		Content: "Ministers and backbenchers will be at loggerheads up until the next election, one minister predicts.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-67705549?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "What is GDP and how is it measured?",
		Content: "A basic guide to how the health of the economy is measured, and why that calculation matters.",
		Link:    "https://www.bbc.co.uk/news/business-13200758?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "What is the UK's plan to send asylum seekers to Rwanda?",
		Content: "The UK government wants to send some asylum seekers to Rwanda but has faced several legal hurdles.",
		Link:    "https://www.bbc.co.uk/news/explainers-61782866?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Can the new Rwanda bill work and what could stop it?",
		Content: "The legislation could set up a politically explosive fight with the courts.",
		Link:    "https://www.bbc.co.uk/news/uk-67643900?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Covid inquiry: What is it investigating and how does it work?",
		Content: "The second round of public hearings examined how ministers made decisions during the pandemic.",
		Link:    "https://www.bbc.co.uk/news/explainers-57085964?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "What the Autumn Statement means for you and your money",
		Content: "Now the speech has been delivered, what does it mean for your finances in the months ahead?",
		Link:    "https://www.bbc.co.uk/news/business-67484095?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Autumn Statement 2023: National Insurance and more key announcements by Jeremy Hunt",
		Content: "Here's what you need to know about Chancellor Jeremy Hunt's tax and spending plans for the year ahead.",
		Link:    "https://www.bbc.co.uk/news/business-67276717?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Benefits: Who gets them and how much do they cost?",
		Content: "The chancellor will confirm how much benefits will go up in the Autumn Statement.",
		Link:    "https://www.bbc.co.uk/news/explainers-63129705?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Who is Jeremy Hunt? The chancellor in charge of nation's finances",
		Content: "The chief financial minister has been in control of public spending at a time of economic turmoil.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-63259600?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "How much money does the UK government raise and spend each year?",
		Content: "A breakdown of where the government's money comes from and where it goes.",
		Link:    "https://www.bbc.co.uk/news/business-45814459?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Who is Nigel Farage? A quick guide to I'm a Celeb contestant",
		Content: "He calls himself a \"hero to some, a villain to others\", but who is the politician turned pundit?",
		Link:    "https://www.bbc.co.uk/news/newsbeat-67477838?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Rishi Sunak’s reshuffle: Who is in the prime minister’s cabinet?",
		Content: "A rundown of the people in Prime Minister Rishi Sunak's cabinet, with new faces, as well as those who have changed roles.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-63376560?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Who is Suella Braverman?",
		Content: "She has lost her job as home secretary after accusing the Met Police of bias in the policing of protests.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-67137926?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Where are asylum seekers being housed in hotels in the UK?",
		Content: "The government wants to end contracts with hotels for housing asylum seekers, in order to cut costs.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-67206459?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Keir Starmer: Labour leader hoping for keys to Downing Street",
		Content: "A profile of the former lawyer hoping to return his party to power after 13 years in opposition.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-66304053?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Rishi Sunak: The Star Wars fan turned political force",
		Content: "After a year in office, the prime minister is attempting to change his image ahead of an election.",
		Link:    "https://www.bbc.co.uk/news/business-51490893?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Liberal Democrats: Who is leader Sir Ed Davey?",
		Content: "A profile of the former cabinet minister who once won an award for a lifesaving act of bravery.",
		Link:    "https://www.bbc.co.uk/news/uk-politics-53888106?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "What is net zero and how are the UK and other countries doing?",
		Content: "Carbon emissions need to be cut to \"net zero\" by 2050 to avoid the most dangerous climate impacts.",
		Link:    "https://www.bbc.co.uk/news/science-environment-58874518?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Donald Trump's failed immunity appeal is still a win for his delay strategy",
		Content: "The time it took to issue the decision has indefinitely delayed Mr Trump's election subversion trial.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68061183?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Michigan school shooter's mother guilty of involuntary manslaughter",
		Content: "Jennifer Crumbley is the first US parent ever convicted of manslaughter over a mass shooting by their child.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68223118?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Toby Keith: Country music legend dies aged 62",
		Content: "Toby Keith, one of country music's biggest stars, had previously revealed he had stomach cancer.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68215639?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Why did Harry and Meghan leave the Royal Family and where do they get their money?",
		Content: "The Duke and Duchess of Sussex are not funded by the Royal Family and have various commercial deals.",
		Link:    "https://www.bbc.co.uk/news/explainers-51047186?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Boeing safety checks to use 'more boots on ground'",
		Content: "The head of the US aviation regulator told US lawmakers they were stepping up safety checks.",
		Link:    "https://www.bbc.co.uk/news/business-68220625?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Storm-battered California told to brace for more rain",
		Content: "A powerful storm that battered the state over the weekend is expected to bring more rain on Tuesday.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68206819?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Police to pay $1.9m to black family held at gunpoint in Colorado",
		Content: "The 2020 stop in Colorado was recorded by bystanders, which showed children handcuffed and crying.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68211285?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Super Bowl 58: Taylor Swift conspiracy theories 'nonsense' - NFL boss Roger Goodell",
		Content: "Super Bowl conspiracy theories involving Taylor Swift are \"nonsense\", says NFL commissioner Roger Goodell.",
		Link:    "https://www.bbc.co.uk/sport/american-football/68206676?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Nikki Haley seeks Secret Service protection, citing 'multiple issues'",
		Content: "The last challenger to Donald Trump says her campaign has had \"multiple issues\" with security.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68213181?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Los Angeles: Helicopter crew rescues man from surging river",
		Content: "The man entered the water to save his dog after it was swept away. The dog managed to swim to safety.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68214577?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Killer Mike dismisses arrest at Grammys as 'speed bump'",
		Content: "The 48-year-old was detained on a misdemeanour charge after winning three awards in the rap category.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68201021?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "US Senate releases deal on border and Ukraine - but will it ever become law?",
		Content: "The bill significantly overhauls border and asylum processes, but House Republicans have rejected it.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68143887?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Navigating cancer takes hope, Biden says after King Charles diagnosis",
		Content: "Buckingham Palace announced on Monday that the King was being treated for an unspecified type of cancer.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68169883?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Police apologise to alleged victim in ice hockey sexual assault case",
		Content: "It took years for Canadian police to lay charges against five current or ex-NHL players accused of sexual assault.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68211284?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "What is the Pineapple Express and why has it drenched California?",
		Content: "The record rain pounding California is part of a common weather pattern.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68218352?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Nevada caucuses or primary: Why both Trump and Haley may claim victory",
		Content: "Republicans are holding two separate ballots this week, with the candidates not going head to head.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68208325?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "'Abnormal is normal' when it comes to California's weather",
		Content: "What is a pineapple express and what are the pros and cons of the storms bracing California?",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68212194?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "The highs and lows of Grammys 2024 - and why Taylor Swift won album of year",
		Content: "The best (and worst) performances, the memorable moments, and why Swift's big win was inevitable.",
		Link:    "https://www.bbc.co.uk/news/entertainment-arts-68200931?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Red carpet and ceremony at Grammy Awards in pictures",
		Content: "See the most eye-catching images and outfits as stars attend the music industry's biggest party.",
		Link:    "https://www.bbc.co.uk/news/entertainment-arts-68201194?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Why did US wait to retaliate for drone attack on its troops?",
		Content: "Foreign policy experts said the approach may be calculated to avoid escalating a conflict with Iran.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68182658?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Should we fear an attack of the voice clones?",
		Content: "Audio deepfakes are easy to make, hard to detect, and getting more convincing, experts say.",
		Link:    "https://www.bbc.co.uk/news/technology-68074257?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Misinformation spreads in China on ‘civil war’ in Texas",
		Content: "Chinese social media cheers intensifying standoff between Texas and the White House over illegal migrants.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68185317?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "How Georgia prosecutor affair affects a Trump trial",
		Content: "The Georgia district attorney acknowledged a relationship with Nathan Wade but denies conflict of interest claims.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68038189?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Scientists in awe at detail in telescope photos",
		Content: "",
		Link:    "https://www.bbc.co.uk/news/articles/cz9e7nv0d2yo?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Drone video shows graffiti covering luxury high-rise in LA",
		Content: "Police say 27 floors were painted after over a dozen people entered the abandoned complex last week.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68209951?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Ros Atkins on… Iran, its proxies and the ‘Axis of Resistance’",
		Content: "The BBC's analysis editor explains the Iran-backed militia groups at the centre of US strikes.",
		Link:    "https://www.bbc.co.uk/news/world-68176257?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Groundhog Day: Did Punxsutawney Phil see his shadow?",
		Content: "America's favourite groundhog - and meteorologist - predicted an early spring at this year's ceremony.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68172216?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "‘They’re climbing on my car’: Alaskan bears spotted in Florida",
		Content: "Two kodiak bear cubs were spotted wandering more than 3,600 miles (5,800 km) from their usual habitat.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68177127?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Jennifer Crumbley takes stand at involuntary manslaughter trial",
		Content: "The mother of Ethan Crumbley said she was unaware he had serious mental health issues despite warnings, ahead of the 2021 school shooting.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68176413?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Lloyd Austin: 'I did not handle this right'",
		Content: "The US defence secretary said he should have informed the president and public of his cancer diagnosis.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68172214?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Why these Arab Americans say Biden has lost their vote",
		Content: "The US president's support for Israel has left some Arab American voters feeling alienated.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68172215?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Sea of foam floods Houston airport hangar",
		Content: "The fire suppression foam was released by accident, filling the building and spilling out onto the surrounding road.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-68176255?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Canada 'Sixties Scoop': Indigenous survivors map out their stories",
		Content: "Canada's \"Sixties Scoop\" saw thousands of indigenous children forcibly removed from their families.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-55269251?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "New Christmas campaign for Canadians held in China",
		Content: "The campaign is based on efforts around a similar case involving a British journalist fifty years ago.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-55249770?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "New Brunswick outbreak: How a smalltown doctor became a Covid pariah",
		Content: "After being labelled \"patient zero\", a small-town physician in Canada was shunned.",
		Link:    "https://www.bbc.co.uk/news/world-us-canada-54686672?at_medium=RSS&at_campaign=KARANGA",
	},

	{
		Title:   "Some on-air claims about Dominion Voting Systems were false, Fox News acknowledges in statement after deal is announced",
		Content: "",
		Link:    "https://www.cnn.com/business/live-news/fox-news-dominion-trial-04-18-23/index.html",
	},

	{
		Title:   "Dominion still has pending lawsuits against election deniers such as Rudy Giuliani and Sidney Powell",
		Content: "",
		Link:    "https://www.cnn.com/business/live-news/fox-news-dominion-trial-04-18-23/h_8d51e3ae2714edaa0dace837305d03b8",
	},

	{
		Title:   "Here are the 20 specific Fox broadcasts and tweets Dominion says were defamatory",
		Content: "• Fox-Dominion trial delay 'is not unusual,' judge says\n• Fox News' defamation battle isn't stopping Trump's election lies",
		Link:    "https://www.cnn.com/2023/04/17/media/dominion-fox-news-allegations/index.html",
	},

	{
		Title:   "Judge in Fox News-Dominion defamation trial: 'The parties have resolved their case'",
		Content: "The judge just announced in court that a settlement has been reached in the historic defamation case between Fox News and Dominion Voting Systems.",
		Link:    "https://www.cnn.com/2023/04/18/media/fox-dominion-settlement/index.html",
	},

	{
		Title:   "'Difficult to say with a straight face': Tapper reacts to Fox News' statement on settlement",
		Content: "A settlement has been reached in Dominion Voting Systems' defamation case against Fox News, the judge for the case announced. The network will pay more than $787 million to Dominion, a lawyer for the company said.",
		Link:    "https://www.cnn.com/videos/politics/2023/04/18/jake-tapper-dominion-lawsuit-settlement-fox-news-statement-lead-vpx.cnn",
	},

	{
		Title:   "Millions in the US could face massive consequences unless McCarthy can navigate out of a debt trap he set for Biden",
		Content: "• DeSantis goes to Washington, a place he once despised, looking for support to take on Trump\n• Opinion: For the GOP to win, it must ditch Trump\n• Chris Christie mulling 2024 White House bid\n• Analysis: The fire next time has begun burning in Tennessee ",
		Link:    "https://www.cnn.com/2023/04/18/politics/mccarthy-biden-debt-ceiling/index.html",
	},

	{
		Title:   "White homeowner accused of shooting a Black teen who rang his doorbell turns himself in to face criminal charges",
		Content: "• 'A major part of Ralph died': Aunt of teen shot after ringing wrong doorbell speaks\n• 20-year-old woman shot after friend turned into the wrong driveway in upstate New York, officials say",
		Link:    "https://www.cnn.com/2023/04/18/us/kansas-city-ralph-yarl-shooting-tuesday/index.html",
	},

	{
		Title:   "Newly released video shows scene of Jeremy Renner's snowplow accident",
		Content: "Newly released body camera footage shows firefighters and sheriff's deputies rushing to help actor Jeremy Renner after a near-fatal snowplow accident in January. The \"Avengers\" actor broke more than 30 bones and suffered other severe injuries. CNN's Chloe Melas has more. ",
		Link:    "https://www.cnn.com/videos/us/2023/04/18/jeremy-renner-snowplow-accident-bodycam-nc-melas-contd-vpx.cnn",
	},

	{
		Title:   "Jake Gyllenhaal and Jamie Lee Curtis spent the Covid-19 lockdown together",
		Content: "It's sourdough bread and handstands for Jake Gyllenhaal and Jamie Lee Curtis.",
		Link:    "https://www.cnn.com/2023/04/18/entertainment/jake-gyllenhaal-jamie-lee-curtis-pandemic-living/index.html",
	},

	{
		Title:   "Toddler crawls through White House fence, prompts Secret Service response",
		Content: "A tiny intruder infiltrated White House grounds Tuesday, prompting a swift response from the US Secret Service.",
		Link:    "https://www.cnn.com/2023/04/18/politics/white-house-toddler/index.html",
	},

	{
		Title:   "Jamie Foxx remains hospitalized nearly a week after 'medical complication'",
		Content: "Jamie Foxx remains hospitalized in Georgia nearly a week after his daughter revealed the actor experienced a \"medical complication,\" a source with knowledge of the matter told CNN on Monday.",
		Link:    "https://www.cnn.com/2023/04/17/entertainment/jamie-foxx-remains-hospitalized/index.html",
	},

	{
		Title:   "A 13-year-old dies after participating in a Benadryl TikTok 'challenge'",
		Content: "A 13-year-old in Ohio has died after \"he took a bunch of Benadryl,\" trying a dangerous TikTok challenge that's circulating online, according to a CNN affiliate and a GoFundMe account from his family.",
		Link:    "https://www.cnn.com/2023/04/18/us/benadryl-tiktok-challenge-teen-death-wellness/index.html",
	},

	{
		Title:   "See pizza delivery guy take out suspect fleeing police",
		Content: "Pizza guy delivers more than a pie, taking out a fleeing suspect. CNN's Jeanne Moos shows him putting his best foot forward.",
		Link:    "https://www.cnn.com/videos/us/2023/04/18/pizza-guy-trips-perp-moos-cprog-orig-bdk.cnn",
	},

	{
		Title:   "Netflix is winding down its DVD business after 25 years",
		Content: "Netflix is officially winding down the business that helped make it a household name. ",
		Link:    "https://www.cnn.com/2023/04/18/media/netflix-dvd-red-envelopes/index.html",
	},

	{
		Title:   "FTC chair Lina Khan warns AI could 'turbocharge' fraud and scams",
		Content: "Artificial intelligence tools such as ChatGPT could lead to a \"turbocharging\" of consumer harms including fraud and scams, and the US government has substantial authority to crack down on AI-driven consumer harms under existing law, members of the Federal Trade Commission said Tuesday. ",
		Link:    "https://www.cnn.com/2023/04/18/tech/lina-khan-ai-warning/index.html",
	},

	{
		Title:   "Eating too much of these foods is driving the rise in type 2 diabetes, study says",
		Content: "Gobbling up too many refined wheat and rice products, along with eating too few whole grains, is fueling the growth of new cases of type 2 diabetes worldwide, according to a new study that models data through 2018.",
		Link:    "https://www.cnn.com/2023/04/17/health/rise-type-2-diabetes-global-wellness/index.html",
	},

	{
		Title:   "ADHD medication abuse in schools is a 'wake-up call'",
		Content: "At some middle and high schools in the United States, 1 in 4 teens report they've abused prescription stimulants for attention deficit hyperactivity disorder during the year prior, a new study found. ",
		Link:    "https://www.cnn.com/2023/04/18/health/teen-misuse-adhd-meds-wellness/index.html",
	},

	{
		Title:   "Apple CEO was presented with an original Macintosh. See his reaction",
		Content: "CEO Tim Cook personally welcomed customers to the new Apple store in Mumbai as the tech company opens its first retail stores in India. CNN's Vedika Sud reports. ",
		Link:    "https://www.cnn.com/videos/tech/2023/04/18/apple-store-mumbai-india-ceo-tim-cook-vedika-sud-ovn-biz-ldn-vpx.cnn",
	},

	{
		Title:   "Democrats bash Justice Clarence Thomas but their plan to investigate ethics allegations is unclear",
		Content: "Senate Democrats railed against Justice Clarence Thomas on Tuesday amid reports that the Supreme Court conservative failed to disclose luxury travel, gifts and a real estate transaction involving a GOP megadonor, but their plan to investigate the conservative jurist remains unclear.  ",
		Link:    "https://www.cnn.com/2023/04/18/politics/clarence-thomas-ethics-democrats/index.html",
	},

	{
		Title:   "Russia is 'going backwards' in equipment and deploying post WWII-era tanks, according to Western officials",
		Content: "• Jailed Wall Street Journal reporter Evan Gershkovich denied detention appeal in Moscow\n• Putin visits Russian troops at military headquarters in Kherson\n• Watch moment WSJ journalist appears in Russian court",
		Link:    "https://www.cnn.com/europe/live-news/russia-ukraine-war-news-04-18-23/index.html",
	},

	{
		Title:   "Two Russians claiming to be former Wagner commanders admit killing children and civilians in Ukraine",
		Content: "Two Russian men who claim to be former Wagner Group commanders have told a human rights activist that they killed children and civilians during their time in Ukraine. ",
		Link:    "https://www.cnn.com/2023/04/17/europe/wagner-commanders-russia-kill-children-intl-hnk/index.html",
	},

	{
		Title:   "'My stomach is hurting from laughing': Hear panelist's reaction to DeSantis' threat to Disney",
		Content: "CNN panelists react to Florida Gov. Ron DeSantis floating the idea of building a competing theme park next to Disney World in Orlando.",
		Link:    "https://www.cnn.com/videos/business/2023/04/18/desantis-disney-competition-panel-reax-pt-vpx.cnn",
	},

	{
		Title:   "GOP prepared to block vote to replace Feinstein on Senate Judiciary",
		Content: "Senate Majority Leader Chuck Schumer said on Tuesday that he hopes to replace Democratic Sen. Dianne Feinstein on the Senate Judiciary Committee with Sen. Ben Cardin of Maryland and aims to set up a floor vote on the issue this afternoon, which Republicans are expected to block. ",
		Link:    "https://www.cnn.com/2023/04/18/politics/schumer-senate-feinstein-vote-cardin/index.html",
	},

	{
		Title:   "Oklahoma governor calls on officials to resign over recording of racist and threatening remarks",
		Content: "",
		Link:    "https://www.cnn.com/2023/04/18/us/mccurtain-county-oklahoma-officials-recording/index.html",
	},

	{
		Title:   "McCarthy slams Biden in handling of US debt ",
		Content: "House Speaker Kevin McCarthy traveled to Wall Street on Monday to deliver a fresh warning that the House GOP majority will refuse to lift a cap on government borrowing unless Biden agrees to spending cuts that would effectively neutralize his domestic agenda.",
		Link:    "https://www.cnn.com/videos/politics/2023/04/18/kevin-mccarthy-wall-street-speech-debt-ceiling-biden-economy-vpx.cnn",
	},

	{
		Title:   "US warns Russia not to touch American nuclear technology at Ukrainian nuclear plant",
		Content: "The US has sensitive nuclear technology at a nuclear power plant inside Ukraine and is warning Russia not to touch it, according to a letter the US Department of Energy sent to Russia's state-owned nuclear energy firm Rosatom last month.",
		Link:    "https://www.cnn.com/2023/04/18/politics/us-warns-russia-zaporizhzhia-nuclear-plant/index.html",
	},

	{
		Title:   "Repeated gunshots fired on live TV as ex-lawmaker shot by assassins",
		Content: "Atiq Ahmed, a former lawmaker in India's parliament, convicted of kidnapping, was shot dead along with his brother while police were escorting them for a medical check-up in a slaying caught on live television on Saturday. CNN's Vedika Sud reports.",
		Link:    "https://www.cnn.com/videos/world/2023/04/18/india-ex-lawmaker-atiq-ahmed-assassination-sud-pkg-contd-ovn-intl-hnk-vpx.cnn",
	},

	{
		Title:   "FDA clears the way for additional bivalent boosters for certain vulnerable individuals",
		Content: "The U.S. Food and Drug Administration amended the terms of its emergency use authorizations for the Pfizer and Moderna bivalent vaccines on Tuesday, allowing people ages 65 and older and certain people with weakened immunity to get additional doses before this fall's vaccination campaigns.",
		Link:    "https://www.cnn.com/2023/04/18/health/fda-bivalent-booster-additional-doses/index.html",
	},

	{
		Title:   "Maine authorities detained a person of interest after 4 people were found dead in a home and 3 others shot while driving",
		Content: "Maine authorities have detained a person of interest and continue to investigate after two shooting incidents that appear to be connected left at least four people dead and three others injured, state police said. ",
		Link:    "https://www.cnn.com/2023/04/18/us/maine-shooting-bowdoin-yarmouth/index.html",
	},

	{
		Title:   "Southwest says flights resumed after delays caused by 'tech issues'",
		Content: "• Delta Air Lines reports record bookings for summer travel\n• Air France and Airbus acquitted in trial over 2009 plane crash",
		Link:    "https://www.cnn.com/travel/article/southwest-airlines-flight-delays/index.html",
	},

	{
		Title:   "Damar Hamlin cleared to resume football activities after January cardiac arrest",
		Content: "Buffalo Bills safety Damar Hamlin, who has been cleared to resume football activities, said Tuesday his cardiac arrest during an NFL game in January was caused by commotio cordis.",
		Link:    "https://www.cnn.com/2023/04/18/sport/damar-hamlin-cleared-to-train-nfl-spt-intl/index.html",
	},

	{
		Title:   "Pilot makes history after landing on top of a 56-story hotel",
		Content: "Polish pilot Lukasz Czepiela made history after landing a plane on a helipad at the top of a 56-story hotel in Dubai.",
		Link:    "https://www.cnn.com/videos/travel/2023/03/17/pilot-lands-on-dubai-helipad-cprog-orig-aw-ao.cnn",
	},

	{
		Title:   "Top US Navy admiral defends non-binary sailor amid some Republican criticism",
		Content: "The top US Navy admiral ardently defended a non-binary sailor on Tuesday amid some criticism from Republican lawmakers, saying he is \"particularly proud of this sailor.\" ",
		Link:    "https://www.cnn.com/2023/04/18/politics/gilday-defends-non-binary-sailor/index.html",
	},

	{
		Title:   "Fulton County DA says fake Trump electors are incriminating one another and wants lawyer disqualified ",
		Content: "The Fulton County District Attorney's office said some fake electors for Donald Trump have implicated each other in potential criminal activity and is seeking to disqualify their lawyer, according to a new court filing.",
		Link:    "https://www.cnn.com/2023/04/18/politics/fulton-county-trump-fake-electors/index.html",
	},

	{
		Title:   "High speed trains are racing across the world. But not in America",
		Content: "High speed trains have proved their worth across the world over the past 50 years. ",
		Link:    "https://www.cnn.com/travel/article/high-speed-rail-us/index.html",
	},

	{
		Title:   "Podcast: One country musician is calling for other artists to oppose assault rifles",
		Content: "",
		Link:    "https://www.cnn.com/audio/podcasts/the-assignment/episodes/42a2f0e2-066a-4675-82ff-afe2016a0bb5",
	},

	{
		Title:   "Here's what you need to know if you haven't filed your return yet — and even if you have",
		Content: "It's April 18, the official deadline to file your federal and state income tax returns for 2022. (It is also, apparently, National Animal Crackers Day for those who celebrate.)",
		Link:    "https://www.cnn.com/2023/04/18/success/tax-day-2023-file-irs/index.html",
	},

	{
		Title:   "Undocumented immigrants are paying their taxes today, too",
		Content: "It's a surprising fact that's often overlooked in the immigration debate. ",
		Link:    "https://www.cnn.com/2023/04/18/us/undocumented-immigrants-taxes-cec/index.html",
	},

	{
		Title:   "Opinion: Why millionaires like us want to pay more in taxes",
		Content: "Tuesday is Tax Day in America, one of the most stressful days of the year, when many taxpayers will finally end their procrastination, file their federal returns, and hope for a refund from the IRS. But for many of the nation's wealthiest, it's just another Tuesday. ",
		Link:    "https://www.cnn.com/2023/04/17/opinions/us-tax-system-wealthy-disney-pearl/index.html",
	},

	{
		Title:   "'World's longest' purpose-built cycling tunnel opens in Norway",
		Content: "There are many ways to explore the seven mountains that surround the picturesque UNESCO World Heritage city of Bergen on Norway's fjord-studded west coast. The newest, however, might well be record-breaking. ",
		Link:    "https://www.cnn.com/travel/article/worlds-longest-cycling-tunnel/index.html",
	},

	{
		Title:   "Artist rejects photo prize after AI-generated image wins award",
		Content: "A German artist has rejected an award from a prestigious international photography competition after revealing that his submission was generated by Artificial Intelligence (AI). ",
		Link:    "https://www.cnn.com/style/article/ai-photo-win-sony-scli-intl/index.html",
	},

	{
		Title:   "These ships disappeared in Lake Superior a century ago. Watch as they're found again",
		Content: "The Great Lakes Shipwreck Historical Society has found two of three ships that sank in the same Lake Superior storm more than a century ago, locating one in 2021 and the other in 2022.",
		Link:    "https://www.cnn.com/videos/travel/2023/04/18/century-old-shipwrecks-lake-superior-discovery-contd-orig-zt.cnn",
	},

	{
		Title:   "Erotic images of seniors show sex and intimacy in new light",
		Content: "What does intimacy look like for seniors? There's no end to sex scenes and other steamy content featuring the young and unwrinkled, but past a certain age, popular culture largely draws a blank — or treats sex as a punchline. ",
		Link:    "https://www.cnn.com/style/article/marilyn-minter-artist-elder-sex/index.html",
	},

	{
		Title:   "China's economy is off to a solid start, rising 4.5% in Q1 2023 ",
		Content: "China's economy is off to a solid start in 2023 following its emergence from three years of strict pandemic restrictions.",
		Link:    "https://www.cnn.com/2023/04/17/economy/china-gdp-q1-2023-intl-hnk/index.html",
	},

	{
		Title:   "Even when wives make as much as husbands, they still do more at home",
		Content: "• Four out of the five US metro areas with the lowest unemployment are in Florida. Here's why\n• Opinion: The overlooked problem with raising the retirement age for Social Security",
		Link:    "https://www.cnn.com/2023/04/16/success/husbands-wives-earning-division-of-labor-pew-survey/index.html",
	},

	{
		Title:   "McDonald's is upgrading its burgers",
		Content: "McDonald's, which has been focusing on upgrading its core items to boost sales, is rolling out a series of changes designed to improve its signature burgers. ",
		Link:    "https://www.cnn.com/2023/04/17/business/mcdonalds-burgers/index.html",
	},

	{
		Title:   "Google-parent stock drops on fears it could lose search market share to AI-powered rivals",
		Content: "Shares of Google-parent Alphabet fell more than 3% in early trading Monday after a report sparked concerns that its core search engine could lose market share to AI-powered rivals, including Microsoft's Bing.",
		Link:    "https://www.cnn.com/2023/04/17/tech/google-ai-search-engine-stock-drop/index.html",
	},

	{
		Title:   "Bidets save you money and reduce waste — we tested the best options out there",
		Content: "",
		Link:    "https://www.cnn.com/cnn-underscored/reviews/best-bidets?iid=CNNUnderscoredHPcontainer",
	},

	{
		Title:   "50+ products to make your life easier and our planet cleaner",
		Content: "",
		Link:    "https://www.cnn.com/cnn-underscored/home/editors-favorite-sustainable-products?iid=CNNUnderscoredHPcontainer",
	},

	{
		Title:   "Mother's Day is around the corner. Here are 50+ thoughtful gifts she'll love",
		Content: "",
		Link:    "https://www.cnn.com/cnn-underscored/gifts/best-mothers-day-gifts-2023?iid=CNNUnderscoredHPcontainer",
	},

	{
		Title:   "A head-to-toe guide of how men should dress this spring, and where they should shop",
		Content: "",
		Link:    "https://www.cnn.com/cnn-underscored/fashion/mens-spring-fashion-style-guide?iid=CNNUnderscoredHPcontainer",
	},

	{
		Title:   "42 of the most useful travel products you can buy on Amazon",
		Content: "",
		Link:    "https://www.cnn.com/cnn-underscored/travel/amazon-travel-products?iid=CNNUnderscoredHPcontainer",
	},

	{
		Title:   "The 7 best high-yield savings accounts of April 2023",
		Content: "",
		Link:    "https://www.cnn.com/cnn-underscored/money/high-yield-savings-accounts?iid=CNNUnderscoredHPcontainer",
	},

	{
		Title:   "Taxes are due tomorrow. Here's how to file for an extension",
		Content: "",
		Link:    "https://www.cnn.com/cnn-underscored/money/how-to-file-taxes?iid=CNNUnderscoredHPcontainer",
	},

	{
		Title:   "Composting is an easy way to reduce food waste. Here's how to do it",
		Content: "",
		Link:    "https://www.cnn.com/cnn-underscored/home/how-to-compost-at-home?iid=CNNUnderscoredHPcontainer",
	},

	{
		Title:   "We stopped using aluminum foil for cooking and you should too. Here's what to use instead",
		Content: "",
		Link:    "https://www.cnn.com/cnn-underscored/reviews/mmmat-silicone-mats?iid=CNNUnderscoredHPcontainer",
	},

	{
		Title:   "The beloved Dyson Supersonic hair dryer is at its lowest price ever",
		Content: "",
		Link:    "https://www.cnn.com/cnn-underscored/deals/dyson-supersonic-sale-2023-04-17?iid=CNNUnderscoredHPcontainer?iid=CNNUnderscoredHPcontainer",
	},

	{
		Title:   "Everything you need to know about Way Day 2023, Wayfair's biggest sale of the year",
		Content: "",
		Link:    "https://www.cnn.com/cnn-underscored/deals/wayfair-way-day-2023-04-17?iid=CNNUnderscoredHPcontainer",
	},

	{
		Title:   "The 10 best Amazon deals to shop this week",
		Content: "",
		Link:    "https://www.cnn.com/cnn-underscored/deals/best-amazon-deals-2023-04-12?iid=CNNUnderscoredHPcontainer",
	},

	{
		Title:   "Mifepristone saved my life ",
		Content: "The ruling earlier this month by a Texas federal judge to suspend the US Food and Drug Administration's approval of a drug that is used frequently for medication abortions, is very personal for me.  ",
		Link:    "https://www.cnn.com/2023/04/18/opinions/medication-abortion-mifepristone-miscarriage-jones-ctpr/index.html",
	},

	{
		Title:   "The 2024 presidential alternative many voters will want",
		Content: "",
		Link:    "https://www.cnn.com/2023/04/18/opinions/2024-presidential-election-alternative-voters-lieberman",
	},

	{
		Title:   "Why isn't the House Judiciary Committee looking into Thomas?",
		Content: "On Monday, the GOP-controlled House Judiciary Committee — chaired by Donald Trump ally Rep. Jim Jordan — is set to hold a field hearing in New York City called \"Victims of Violent Crime in Manhattan.\" A statement bills the hearing as an examination of how, the Judiciary Committee says, Manhattan District Attorney Alvin Bragg's policies have \"led to an increase in violent crime and a dangerous community for New York City residents.\" ",
		Link:    "https://www.cnn.com/2023/04/17/opinions/jim-jordan-clarence-thomas-judiciary-committee-obeidallah/index.html",
	},

	{
		Title:   "Top secrets come spilling out",
		Content: "In 1917, British analysts deciphered a coded message the German foreign minister sent to one of his country's diplomats vowing to begin \"unrestricted submarine warfare\" and seeking to win over Mexico with a promise to \"reconquer the lost territory in Texas, New Mexico and Arizona\" if the US entered the world war. When it became public, the Zimmerman Telegram caused a sensation, helping propel the US into the conflict against Germany. ",
		Link:    "https://www.cnn.com/2023/04/16/opinions/top-secrets-come-spilling-out-opinion-column-galant/index.html",
	},

	{
		Title:   "How did Sudan go from casting off despotic rule to this?",
		Content: "Four years ago, almost to the day, the people of Sudan were celebrating a revolution after overthrowing longtime dictator Omar al-Bashir. Now the East African country faces the possibility of a complete collapse similar to the chaos we see today in Yemen or Libya.",
		Link:    "https://www.cnn.com/2023/04/17/opinions/sudan-revolution-to-civil-war-lynch/index.html",
	},

	{
		Title:   "Michelle Yeoh set to return in new 'Star Trek' movie",
		Content: "Live long and prosper, Michelle Yeoh. ",
		Link:    "https://www.cnn.com/2023/04/18/entertainment/michelle-yeoh-star-trek-section-31/index.html",
	},

	{
		Title:   "Recap: 'Succession' finds dark humor in the aftershocks",
		Content: "After the shock came the aftershocks, the power vacuum, and perhaps most significantly and impressively, the laughs, as \"Succession\" pivoted to face life after Logan Roy, in an episode that finally put the HBO show's title into full flower.",
		Link:    "https://www.cnn.com/2023/04/16/entertainment/succession-season-4-episode-4-recap/index.html",
	},

	{
		Title:   "'Yellowjackets' leans hard into '90s music nostalgia, and we're here for it",
		Content: "Of the many dark gifts Showtime's eerie hit series \"Yellowjackets\" serves up for us, the juiciest this season is by far the music.",
		Link:    "https://www.cnn.com/2023/04/14/entertainment/yellowjackets-90s-music/index.html",
	},

	{
		Title:   "Jeremy Renner revisits 'the amazing group of people' who helped him recover from his accident",
		Content: "Jeremy Renner is continuing his recovery after his devastating snowplow accident in January, and recognizing those who've helped him along the way.",
		Link:    "https://www.cnn.com/2023/04/16/entertainment/jeremy-renner-update/index.html",
	},

	{
		Title:   "Review: 'Barry' takes a whack at its farewell season",
		Content: "\"Barry\" has taken chances from the very beginning, which is certainly true of a fourth and final season that picks up where the third left off, with its hitman-turned-wannabe actor getting arrested. That paves the way for an even darker season that accentuates the show's ensemble aspect while leaning a little too heavily on blurring lines with flights of fancy.",
		Link:    "https://www.cnn.com/2023/04/14/entertainment/barry-season-4-review/index.html",
	},

	{
		Title:   "Markets digest bank earnings after recent turmoil",
		Content: "",
		Link:    "https://www.cnn.com/business/live-news/stock-market-bank-earnings/index.html",
	},

	{
		Title:   "Still haven't filed your taxes? Here's what you need to know",
		Content: "So far this tax season, the IRS has received more than 90 million income tax returns for 2022. ",
		Link:    "https://www.cnn.com/2023/04/13/success/tax-filing-tips/index.html",
	},

	{
		Title:   "Retail spending fell in March as consumers pull back",
		Content: "Spending at US retailers fell in March as consumers pulled back amid recessionary fears fueled by the banking crisis.",
		Link:    "https://www.cnn.com/2023/04/14/economy/march-retail-sales/index.html",
	},

	{
		Title:   "Analysis: Fox News is about to enter the true No Spin Zone",
		Content: "This is it. ",
		Link:    "https://www.cnn.com/2023/04/14/media/fox-news-dominion-hnk-intl/index.html",
	},

	{
		Title:   "Silicon Valley Bank collapse renews calls to address disparities impacting entrepreneurs of color ",
		Content: "When customers at Silicon Valley Bank rushed to withdraw billions of dollars last month, venture capitalist Arlan Hamilton stepped in to help some of the founders of color who panicked about losing access to payroll funds.",
		Link:    "https://www.cnn.com/2023/04/13/business/silicon-valley-bank-entrepreneurs-of-color-reaj/index.html",
	},

	{
		Title:   "Not only is Lake Powell's water level plummeting because of drought, its total capacity is shrinking, too",
		Content: "Lake Powell, the second-largest human-made reservoir in the US, has lost nearly 7% of its potential storage capacity since 1963, when Glen Canyon Dam was built, a new report shows. ",
		Link:    "https://www.cnn.com/2022/03/21/us/lake-powell-capacity-shrinking-drought-climate/index.html",
	},

	{
		Title:   "These were the best and worst places for air quality in 2021, new report shows",
		Content: "Air pollution spiked to unhealthy levels around the world in 2021, according to a new report.",
		Link:    "https://www.cnn.com/2022/03/22/world/air-pollution-2021-iqair-report-climate/index.html",
	},

	{
		Title:   "Big-box stores could help slash emissions and save millions by putting solar panels on roofs. Why aren't more of them doing it?",
		Content: "As the US attempts to wean itself off its heavy reliance on fossil fuels and shift to cleaner energy sources, many experts are eyeing a promising solution: your neighborhood big-box stores and shopping malls.",
		Link:    "https://www.cnn.com/2022/03/20/us/solar-power-on-big-box-store-rooftops-climate/index.html",
	},

	{
		Title:   "Look of the Week: Blackpink headline Coachella in Korean hanboks",
		Content: "Bringing the second day of this year's Coachella to a close, K-Pop girl group Blackpink made history Saturday night when they became the first Asian act to ever headline the festival. To a crowd of, reportedly, over 125,000 people, Jennie, Jisoo, Lisa and Rosé used the ground-breaking moment to pay homage to Korean heritage by arriving onstage in hanboks: a traditional type of dress.",
		Link:    "https://www.cnn.com/style/article/blackpink-coachella-2023-hanboks-lotw/index.html",
	},

	{
		Title:   "Scientists identify secret ingredient in Leonardo da Vinci paintings",
		Content: "\"Old Masters\" such as Leonardo da Vinci, Sandro Botticelli and Rembrandt may have used proteins, especially egg yolk, in their oil paintings, according to a new study.",
		Link:    "https://www.cnn.com/style/article/old-masters-da-vinci-egg-yolk-painting-scn/index.html",
	},

	{
		Title:   "How Playboy cut ties with Hugh Hefner to create a post-MeToo brand",
		Content: "Hugh Hefner launched Playboy Magazine 70 years ago this year. The first issue included a nude photograph of Marilyn Monroe, which he had purchased and published without her knowledge or consent.",
		Link:    "https://www.cnn.com/style/article/playboy-the-conversation/index.html",
	},

	{
		Title:   "'A definitive backslide.' Inside fashion's worrying runway trend",
		Content: "Now that the Fall-Winter 2023 catwalks have been disassembled, it's clear one trend was more pervasive than any collective penchant for ruffles, pleated skirts or tailored coats. ",
		Link:    "https://www.cnn.com/style/article/fashion-week-fall-winter-2023-size-diversity-skinny-wegovy/index.html",
	},

	{
		Title:   "Michael Jordan's 1998 NBA Finals sneakers sell for a record $2.2 million",
		Content: "In 1998, Michael Jordan laced up a pair of his iconic black and red Air Jordan 13s to bring home a Bulls victory during Game 2 of his final NBA championship — and now they are the most expensive sneakers ever to sell at auction.\n\nThe game-winning sneakers sold for $2.2 million at Sotheby's in New York on Tuesday, smashing the sneaker auction record of $1.47 million, set in 2021 by a pair of Nike Air Ships that Jordan wore earlier in his career.",
		Link:    "https://www.cnn.com/style/article/michael-jordan-sneakers-1998-finals-sothebys-auction-record/index.html",
	},

	{
		Title:   "The surreal facades of America's strip clubs",
		Content: "Some people travel the world in search of adventure, while others seek out natural wonders, cultural landmarks or culinary experiences. But French photographer François Prost was looking for something altogether different during his recent road trip across America: strip clubs.",
		Link:    "https://www.cnn.com/style/article/francois-prost-gentlemens-club/index.html",
	},

	{
		Title:   "Here's the real reason to turn on airplane mode when you fly",
		Content: "We all know the routine by heart: \"Please ensure your seats are in the upright position, tray tables stowed, window shades are up, laptops are stored in the overhead bins and electronic devices are set to flight mode.\"",
		Link:    "https://www.cnn.com/travel/article/airplane-mode-reasons-why/index.html",
	},

	{
		Title:   "'I was up to my waist down a hippo's throat.' He survived, and here's his advice",
		Content: "Paul Templer was living his best life.",
		Link:    "https://www.cnn.com/travel/article/hippo-attack-avoid-survive-paul-templer/index.html",
	},

	{
		Title:   "They bought an abandoned 'ghost house' in the Japanese countryside",
		Content: "He'd spent years backpacking around the world, and Japanese traveler Daisuke Kajiyama was finally ready to return home to pursue his long-held dream of opening up a guesthouse.",
		Link:    "https://www.cnn.com/travel/article/couple-turned-abandoned-japanese-home-into-guest-house/index.html",
	},

	{
		Title:   "Relaxed entry rules make it easier than ever to visit this stunning Asian nation",
		Content: "Due to its remoteness and short summer season, Mongolia has long been a destination overlooked by travelers. ",
		Link:    "https://www.cnn.com/travel/article/mongolia-reasons-to-visit-2023/index.html",
	},

	{
		Title:   "The most beautiful sections of China's Great Wall",
		Content: "Having lived in Beijing for almost 12 years, I've had plenty of time to travel widely in China. ",
		Link:    "https://www.cnn.com/travel/article/china-beautiful-great-wall-sections-cmd/index.html",
	},

	{
		Title:   "Sign up to our newsletter for a weekly roundup of travel news",
		Content: "",
		Link:    "https://www.cnn.com/specials/travel-newsletter",
	},

	{
		Title:   "Nelly Cheboi, who creates computer labs for Kenyan schoolchildren, is CNN's Hero of the Year",
		Content: "Celebrities and musicians are coming together tonight to honor everyday people making the world a better place.",
		Link:    "https://www.cnn.com/2022/12/11/us/cnn-heroes-all-star-tribute-hero-of-the-year/index.html",
	},

	{
		Title:   "CNN Heroes: Sharing the Spotlight",
		Content: "",
		Link:    "https://www.cnn.com/specials/cnn-heroes-salutes-special",
	},

	{
		Title:   "Donate now to a Top 10 CNN Hero",
		Content: "Anderson Cooper explains how you can easily donate to any of the 2021 Top 10 CNN Heroes.",
		Link:    "https://www.cnn.com/videos/tv/2021/11/26/how-to-donate-matching-cnnheroes.cnn",
	},

	{
		Title:   "0% intro APR until 2024 is 100% insane",
		Content: "",
		Link:    "https://www.fool.com/the-ascent/credit-cards/landing/wells-fargo-reflect-review/?utm_site=theascent&utm_campaign=ta-cc-co-cnn-welref2-ron-5-hp-sfpb&utm_medium=cpc&utm_source=cnn",
	},

	{
		Title:   "It's official: now avoid credit card interest into 2024",
		Content: "",
		Link:    "https://www.fool.com/the-ascent/credit-cards/landing/citi-simplicity-review/?utm_site=theascent&utm_campaign=ta-cc-co-cnn-citisimp2-ron-5-hp-sfpb&utm_medium=cpc&utm_source=cnn",
	},

	{
		Title:   "Experts: this is the best cash back card of 2022",
		Content: "",
		Link:    "https://www.fool.com/the-ascent/credit-cards/landing/wells-fargo-active-cash-card-review/?utm_site=theascent&utm_campaign=ta-cc-co-cnn-welac2-ron-5-hp-sfpb&utm_medium=cpc&utm_source=cnn",
	},

	{
		Title:   "Turn Your Rising Home Equity Into Cash You Can Use",
		Content: "",
		Link:    "https://www.lendingtree.com/?splitterid=home-equity&cproduct=he&cchannel=content&csource=cnn&cmethod=heform&ccreative=risingequitycash_housemoneystack&placement_name=sectionfronts&ad_headline=risingequitycash&ad_image_name=housemoneystack&ctype=sectionfronts&bdst=revshare&mtaid=AC53E&esourceid=6348616",
	},

	{
		Title:   "Dream Big with a Home Equity Loan",
		Content: "",
		Link:    "https://www.lendingtree.com/?splitterid=home-equity&cproduct=he&cchannel=content&csource=cnn&cmethod=heform&ccreative=dreambighomeequity_housemoneystack&placement_name=sectionfronts&ad_headline=dreambighomeequity&ad_image_name=housemoneystack&ctype=sectionfronts&bdst=revshare&mtaid=AC53E&esourceid=6348616",
	},

	{
		Title:   "Want Cash Out of Your Home? Here Are Your Best Options",
		Content: "",
		Link:    "https://www.lendingtree.com/?splitterid=home-equity&cproduct=he&cchannel=content&csource=cnn&cmethod=heform&ccreative=cashoutoptions_housemoneystack&placement_name=sectionfronts&ad_headline=cashoutoptions&ad_image_name=housemoneystack&ctype=sectionfronts&bdst=revshare&mtaid=AC53E&esourceid=6348616",
	},

	{
		Title:   "Fulton County, Georgia, jail leadership resigns after inmate's death and accusations of unsanitary conditions",
		Content: "Three officials at the Fulton County, Georgia, jail have stepped down amid an investigation into the death of an inmate whose family said was housed in a filthy, bug-infested cell that \"was not fit for a diseased animal.\"",
		Link:    "https://www.cnn.com/2023/04/18/us/fulton-county-jail-leadership-resignations-lashawn-thompson-death/index.html",
	},

	{
		Title:   "He was free for 2 years. Now Crosley Green is back in prison for a crime he says he didn't commit",
		Content: "A Florida man who served three decades behind bars for a murder he says he didn't commit returned to prison Monday after spending the past two years building a life outside prison walls. ",
		Link:    "https://www.cnn.com/2023/04/17/us/crosley-green-prison-return/index.html",
	},

	{
		Title:   "Tennessee Air National Guardsman applied to be a hitman online, the FBI says. It was a spoof website and now he's facing charges ",
		Content: "A Tennessee Air National Guardsman is facing charges after applying to be a hitman on a spoof \"rent-a-hitman\" website, according to the Department of Justice.",
		Link:    "https://www.cnn.com/2023/04/17/us/tennessee-air-national-guardsman-hitman-online-application/index.html",
	},

	{
		Title:   "Family of victim and survivors of Indianapolis FedEx mass shooting file lawsuit against gun magazine manufacturer and distributors",
		Content: "The family of a victim and several survivors of a mass shooting at a FedEx facility in Indianapolis filed a lawsuit against companies involved in the manufacturing, marketing and sale of the high capacity magazine used by the gunman who killed 8 people and injured several others in 2021. ",
		Link:    "https://www.cnn.com/2023/04/16/us/lawsuit-indianapolis-fedex-mass-shooting/index.html",
	},

	{
		Title:   "GOP prepared to block vote to replace Feinstein on Senate Judiciary",
		Content: "Senate Majority Leader Chuck Schumer said on Tuesday that he hopes to replace Democratic Sen. Dianne Feinstein on the Senate Judiciary Committee with Sen. Ben Cardin of Maryland and aims to set up a floor vote on the issue this afternoon, which Republicans are expected to block. ",
		Link:    "https://www.cnn.com/2023/04/18/politics/schumer-senate-feinstein-vote-cardin/index.html",
	},

	{
		Title:   "Anti-abortion doctors urge Supreme Court to keep mifepristone restrictions in place",
		Content: "A group of doctors opposed to abortion asked the Supreme Court Tuesday to restrict access to a key medication abortion drug while other legal challenges play out, as Wednesday night's deadline for the court to act approaches.",
		Link:    "https://www.cnn.com/2023/04/18/politics/abortion-supreme-court-mifepristone-restrictions-filing/index.html",
	},

	{
		Title:   "McCarthy makes plea for Republicans to back debt ceiling plan",
		Content: "Speaker Kevin McCarthy made a plea to House Republicans during a closed-door meeting Tuesday morning to back his debt ceiling plan, telling them that although it doesn't have to include everything they want, it will help get him to the negotiating table with President Joe Biden. ",
		Link:    "https://www.cnn.com/2023/04/18/politics/mccarthy-debt-limit-plan/index.html",
	},

	{
		Title:   "Supreme Court seems sympathetic to postal worker who didn't work Sundays in dispute over religious accommodations",
		Content: "The Supreme Court seemed to side with a former mail carrier, an evangelical Christian, who says the US Postal Service failed to accommodate his request to not work on Sundays.",
		Link:    "https://www.cnn.com/2023/04/18/politics/groff-dejoy-supreme-court-religious-liberty/index.html",
	},

	{
		Title:   "China's economy is off to a solid start, rising 4.5% in Q1 2023 ",
		Content: "China's economy is off to a solid start in 2023 following its emergence from three years of strict pandemic restrictions.",
		Link:    "https://www.cnn.com/2023/04/17/economy/china-gdp-q1-2023-intl-hnk/index.html",
	},

	{
		Title:   "Even when wives make as much as husbands, they still do more at home",
		Content: "• Four out of the five US metro areas with the lowest unemployment are in Florida. Here's why\n• Opinion: The overlooked problem with raising the retirement age for Social Security",
		Link:    "https://www.cnn.com/2023/04/16/success/husbands-wives-earning-division-of-labor-pew-survey/index.html",
	},

	{
		Title:   "McDonald's is upgrading its burgers",
		Content: "McDonald's, which has been focusing on upgrading its core items to boost sales, is rolling out a series of changes designed to improve its signature burgers. ",
		Link:    "https://www.cnn.com/2023/04/17/business/mcdonalds-burgers/index.html",
	},

	{
		Title:   "Google-parent stock drops on fears it could lose search market share to AI-powered rivals",
		Content: "Shares of Google-parent Alphabet fell more than 3% in early trading Monday after a report sparked concerns that its core search engine could lose market share to AI-powered rivals, including Microsoft's Bing.",
		Link:    "https://www.cnn.com/2023/04/17/tech/google-ai-search-engine-stock-drop/index.html",
	},

	{
		Title:   "Family throws surprise birthday party at Eagles tailgate for Vietnam vet",
		Content: "",
		Link:    "https://6abc.com/philadelphia-eagles-tailgate-south-vietnam-veteran/12524946/",
	},

	{
		Title:   "$200,000 worth of supplies donated to PB schools",
		Content: "",
		Link:    "https://www.wptv.com/news/education/200-000-worth-of-supplies-distributed-for-palm-beach-county-schools-during-giveaway-event",
	},

	{
		Title:   "Car chase ends with police-involved shooting",
		Content: "",
		Link:    "https://abc7ny.com/police-involved-shooting-grand-concourse-section-suspect-shot-in-head-and-leg-bronx/12524318",
	},

	{
		Title:   "Mother fends off raccoon from daughter",
		Content: "",
		Link:    "https://www.wfsb.com/2022/12/02/racoon-attack-ashford/",
	},

	{
		Title:   "Cars totaling nearly $200k stolen from dealership",
		Content: "",
		Link:    "https://www.atlantanewsfirst.com/2022/12/04/police-2-ford-mustangs-totaling-nearly-200k-stolen-upson-county-dealership/",
	},

	{
		Title:   "Christmas parade road closures and detours",
		Content: "",
		Link:    "https://wach.com/news/local/christmas-parade-road-closures-and-detours-in-downtown-lexington#",
	},

	{
		Title:   "Hospitals becoming a 'dumping ground' for kids",
		Content: "",
		Link:    "https://www.newschannel5.com/news/newschannel-5-investigates/hospitals-becoming-a-dumping-ground-for-kids-in-dcs-custody",
	},

	{
		Title:   "Girl who called 911 for her father honored",
		Content: "",
		Link:    "https://auburnpub.com/news/local/aurelius-girl-who-called-911-for-her-father-honored-by-cayuga-county-sheriff/article_be0826d5-b1ff-51a0-ab94-99dade83ad38.html",
	},

	{
		Title:   "Water study could delay commercial construction",
		Content: "",
		Link:    "https://poststar.com/news/local/lake-george-water-study-could-delay-commercial-construction/article_dc5e67ce-718d-11ed-82c2-ef3febaf7e24.html",
	},

	{
		Title:   "Almost 200 animals rescued from puppy mill ",
		Content: "",
		Link:    "https://www.cbsnews.com/philadelphia/news/almost-200-animals-rescued-from-puppy-mill-in-ocean-county/",
	},

	{
		Title:   "She made a promise to her husband to keep their tree farm going",
		Content: "",
		Link:    "https://www.news5cleveland.com/news/local-news/she-made-a-promise-to-her-husband-to-keep-their-tree-farm-going",
	},

	{
		Title:   "HS football players gain perspective helping vets",
		Content: "",
		Link:    "https://www.kristv.com/news/local-news/high-school-football-players-gain-new-perspective-helping-veterans",
	},

	{
		Title:   "Milwaukee Dancing Grannies planning return",
		Content: "",
		Link:    "https://www.wisn.com/article/milwaukee-dancing-grannies-planning-return-to-waukesha/42138823",
	},

	{
		Title:   "Fire crews respond to fire at boarded up building",
		Content: "",
		Link:    "https://www.kare11.com/article/news/local/mfd-crews-respond-to-fire-at-boarded-up-apartment-building/89-6b440525-7462-4b35-86be-6715c5007f8d",
	},

	{
		Title:   "10-year-old shoot mom over VR headset",
		Content: "",
		Link:    "https://www.tmj4.com/news/local-news/10-year-old-upset-over-vr-headset-fatally-shoots-mother-charged-as-an-adult",
	},

	{
		Title:   "Condo owners remember building scare",
		Content: "",
		Link:    "https://www.cbs58.com/news/horizon-west-condo-owners-in-waukesha-remember-building-fire-one-year-later",
	},

	{
		Title:   "5th grader raises $2,000 for grieving teacher",
		Content: "",
		Link:    "https://www.ketv.com/article/ralston-she-was-going-through-a-lot-5th-grader-raises-2000-for-grieving-teacher/42115993",
	},

	{
		Title:   "How Kyrsten Sinema's decision makes Democrats' 2024 Senate map tighter",
		Content: "Arizona Sen. Kyrsten Sinema decided to shake up the political world on Friday by becoming an independent. The former Democrat is still caucusing with the party in the Senate, so the Democratic caucus still has 51 members. Now, instead of 49 Democrats and two independents within their ranks, the caucus has 48 Democrats and three independents. ",
		Link:    "https://www.cnn.com/2022/12/10/politics/kyrsten-sinema-independent-democrats/index.html",
	},

	{
		Title:   "Meet the woman steering Biden's bipartisan winning streak on Capitol Hill",
		Content: "The Biden administration managed to rack up a long list of major legislative wins in its first two years despite facing one of the most closely-divided Congresses in history.\nFrom bipartisan action on infrastructure, gun safety and same-sex marriage to party-line bills tackling climate change and expanding health care coverage, it's a record President Joe Biden and Democrats on the ballot were all eager to tout on the campaign trail during the midterms. ",
		Link:    "https://www.cnn.com/2022/12/11/politics/louisa-terrell-legislative-director/index.html",
	},

	{
		Title:   "CNN Political Briefing",
		Content: "The political news you need to know, in 10 minutes or less. Hosted by David Chalian.",
		Link:    "https://www.cnn.com/audio/podcasts/political-briefing",
	},

	{
		Title:   "The Axe Files with David Axelrod",
		Content: "Go beyond the soundbites and get to know some of the most interesting players in politics.",
		Link:    "https://www.cnn.com/audio/podcasts/axe-files",
	},

	{
		Title:   "Margins of Error",
		Content: "Look closely at almost anything and you'll find data—lots of it. But what are those numbers really saying about who we are and what we believe? Harry Enten is on a mission to find out.",
		Link:    "https://www.cnn.com/audio/podcasts/margins-of-error",
	},

	{
		Title:   "The notable legal clouds that continue to hang over Donald Trump",
		Content: "All eyes are on former President Donald Trump, who has launched another White House bid. ",
		Link:    "https://www.cnn.com/2022/11/15/politics/donald-trump-investigations-lawsuits/index.html",
	},

	{
		Title:   "The fine print of the Respect for Marriage Act",
		Content: "Let's start with the positive: Republicans and Democrats are coming together to protect same-sex marriage from the Supreme Court. ",
		Link:    "https://www.cnn.com/2022/11/29/politics/respect-for-marriage-act-what-matters/index.html",
	},

	{
		Title:   "Everything you need to know about Biden's student loan forgiveness program",
		Content: "President Joe Biden's federal student loan forgiveness program, which promises to deliver up to $20,000 of debt relief for millions of borrowers, is on hold indefinitely as legal challenges work their way through the courts. ",
		Link:    "https://www.cnn.com/2022/08/31/politics/biden-student-loan-forgiveness-faq/index.html",
	},

	{
		Title:   "Politics of the Day",
		Content: "Politics of the Day",
		Link:    "https://www.cnn.com/video/playlists/this-week-in-politics/",
	},

	{
		Title:   "0% intro APR until 2024 is 100% insane",
		Content: "",
		Link:    "https://www.fool.com/the-ascent/credit-cards/landing/wells-fargo-reflect-review/?utm_site=theascent&utm_campaign=ta-cc-co-cnn-welref2-ron-5-hp-sfpb&utm_medium=cpc&utm_source=cnn",
	},

	{
		Title:   "It's official: now avoid credit card interest into 2024",
		Content: "",
		Link:    "https://www.fool.com/the-ascent/credit-cards/landing/citi-simplicity-review/?utm_site=theascent&utm_campaign=ta-cc-co-cnn-citisimp2-ron-5-hp-sfpb&utm_medium=cpc&utm_source=cnn",
	},

	{
		Title:   "Experts: this is the best cash back card of 2022",
		Content: "",
		Link:    "https://www.fool.com/the-ascent/credit-cards/landing/wells-fargo-active-cash-card-review/?utm_site=theascent&utm_campaign=ta-cc-co-cnn-welac2-ron-5-hp-sfpb&utm_medium=cpc&utm_source=cnn",
	},

	{
		Title:   "Turn Your Rising Home Equity Into Cash You Can Use",
		Content: "",
		Link:    "https://www.lendingtree.com/?splitterid=home-equity&cproduct=he&cchannel=content&csource=cnn&cmethod=heform&ccreative=risingequitycash_housemoneystack&placement_name=sectionfronts&ad_headline=risingequitycash&ad_image_name=housemoneystack&ctype=sectionfronts&bdst=revshare&mtaid=AC53E&esourceid=6348616",
	},

	{
		Title:   "Dream Big with a Home Equity Loan",
		Content: "",
		Link:    "https://www.lendingtree.com/?splitterid=home-equity&cproduct=he&cchannel=content&csource=cnn&cmethod=heform&ccreative=dreambighomeequity_housemoneystack&placement_name=sectionfronts&ad_headline=dreambighomeequity&ad_image_name=housemoneystack&ctype=sectionfronts&bdst=revshare&mtaid=AC53E&esourceid=6348616",
	},

	{
		Title:   "Want Cash Out of Your Home? Here Are Your Best Options",
		Content: "",
		Link:    "https://www.lendingtree.com/?splitterid=home-equity&cproduct=he&cchannel=content&csource=cnn&cmethod=heform&ccreative=cashoutoptions_housemoneystack&placement_name=sectionfronts&ad_headline=cashoutoptions&ad_image_name=housemoneystack&ctype=sectionfronts&bdst=revshare&mtaid=AC53E&esourceid=6348616",
	},

	{
		Title:   "Use the Right Card for Holiday Gifts (Save Hundreds on Interest)",
		Content: "",
		Link:    "https://www.comparecards.com/?splitterid=coca-category-low-interest&mtaid=8631C&esourceid=6497536&utm_source=cnn&tar=sectionfronts&grp=cat-low-interest&utm_content=use+the+right+card+for+holiday+gifts+you+could+save+hundreds&adt=brunetteredcard&placement_name=sectionfronts&ad_headline=use+the+right+card+for+holiday+gifts+you+could+save+hundreds&ad_image_name=brunetteredcard&utm_medium=native&ad_position=0&bdst=revshare",
	},

	{
		Title:   "10 Cards Charging 0% Interest Until 2024",
		Content: "",
		Link:    "https://www.comparecards.com/?splitterid=coca-category-low-interest&mtaid=8631C&esourceid=6497536&utm_source=cnn&tar=sectionfronts&grp=cat-low-interest&utm_content=10+cards+charging+0+interest+until+2024&adt=brunetteredcard&placement_name=sectionfronts&ad_headline=10+cards+charging+0+interest+until+2024&ad_image_name=brunetteredcard&utm_medium=native&ad_position=0&bdst=revshare",
	},

	{
		Title:   "Get a $200 Cash Back Bonus on Holiday Gift Buying",
		Content: "",
		Link:    "https://www.comparecards.com/?splitterid=coca-cat-cash-back&mtaid=8631C&esourceid=6488886&utm_source=cnn&tar=sectionfronts&grp=cat-cb&utm_content=get+a+200+cash+back+bonus+on+holiday+gift+buying&adt=brunetteredcard&placement_name=sectionfronts&ad_headline=get+a+200+cash+back+bonus+on+holiday+gift+buying&ad_image_name=brunetteredcard&utm_medium=native&ad_position=0&bdst=revshare",
	},

	{
		Title:   "GOP-led House to vote on impeaching Homeland Security Secretary Mayorkas",
		Content: "The House is set to vote on whether to impeach Department of Homeland Security Secretary Alejandro Mayorkas. Follow here for the latest live news updates.",
		Link:    "https://www.cnn.com/politics/live-news/border-bill-mayorkas-impeachment-vote/index.html",
	},

	{
		Title:   "",
		Content: "President Joe Biden admonished former President Donald Trump and Republicans who are derailing a long-negotiated border and national security bill that is expected to fail in the Senate. ",
		Link:    "https://www.cnn.com/videos/politics/2024/02/06/joe-biden-donald-trump-border-bill-cnc-vpx.cnn",
	},

	{
		Title:   "",
		Content: "CNN chief national correspondent John King heads to South Carolina ahead of its Republican primary to speak with voters about the race between the state's former governor, Nikki Haley, and former president Donald Trump.",
		Link:    "https://www.cnn.com/videos/politics/2024/02/06/john-king-south-carolina-voters-haley-trump-vpx.cnn",
	},

	{
		Title:   "Court rules Trump does not have immunity from 2020 election subversion prosecution",
		Content: "Donald Trump is not immune from prosecution for alleged crimes he committed during his presidency to reverse the 2020 election results, a federal appeals court said Tuesday.",
		Link:    "https://www.cnn.com/politics/live-news/trump-court-ruling-immunity-election-subversion-prosecution/index.html",
	},

	{
		Title:   "",
		Content: "CNN senior legal analyst Elie Honig breaks down an appeals court ruling that former President Donald Trump does not have blanket presidential immunity. The ruling is a major blow to Trump's defense in the federal election subversion case brought by special counsel Jack Smith. ",
		Link:    "https://www.cnn.com/videos/politics/2024/02/06/trump-immunity-ruling-appeals-court-honig-vpx.cnn",
	},

	{
		Title:   "",
		Content: "Donald Trump is not immune from prosecution for alleged crimes he committed during his presidency to reverse the 2020 election results, a federal appeals court said. CNN's Chief Legal Affairs Correspondent Paula Reid details what next steps may look like.",
		Link:    "https://www.cnn.com/videos/politics/2024/02/06/trump-presidential-immunity-federal-appeals-court-ruling-time-line-paula-reid-vpx.cnn",
	},

	{
		Title:   "",
		Content: "Rep. Chip Roy (R-TX) joins CNN's Kaitlan Collins to discuss the Senate's proposed border deal package.",
		Link:    "https://www.cnn.com/videos/politics/2024/02/06/chip-roy-border-bill-src-bts-vpx.cnn",
	},

	{
		Title:   "",
		Content: "Sen. James Lankford (R-OK) responds to Elon Musk's criticism of the border bill that Republicans and Democrats in the Senate negotiated to try and address the surge in migrants at the US-Mexico border. ",
		Link:    "https://www.cnn.com/videos/politics/2024/02/06/border-bill-james-lankford-elon-musk-lead-vpx.cnn",
	},

	{
		Title:   "",
		Content: "During a radio interview on \"The Dan Bongino Show,\" former President Donald Trump said he would like to debate President Joe Biden \"immediately.\" Biden, however, seemed to just laugh the idea off when asked about it by reporters. ",
		Link:    "https://www.cnn.com/videos/politics/2024/02/06/trump-challenges-biden-presidential-debate-cnntm-ldn-vpx.cnn",
	},

	{
		Title:   "",
		Content: "CNN's Abby Phillip breaks down RNC chief Ronna McDaniel's loyalty toward former President Donald Trump despite his public doubts about her future after a series of election-year losses. ",
		Link:    "https://www.cnn.com/videos/politics/2024/02/06/rnc-chief-ronna-mcdaniel-trump-nn-vpx.cnn",
	},

	{
		Title:   "",
		Content: "New York Times reporter and CNN senior political analyst Maggie Haberman joins Jake Tapper to discuss some of the legal woes former President Trump is facing this week.",
		Link:    "https://www.cnn.com/videos/politics/2024/02/05/trump-legal-woes-haberman-tapper-sot-lead-vpx.cnn",
	},

	{
		Title:   "",
		Content: "Rep. Tim Burchett (R-TN) joins The Lead.",
		Link:    "https://www.cnn.com/videos/politics/2024/02/05/rep-tim-burchett-gop-border-bill-no-deal-the-lead-jake-tapper.cnn",
	},

	{
		Title:   "On the Border, Republicans Set a Trap, Then Fell Into It",
		Content: "Republicans are rapidly abandoning a legislative compromise that would have given aid to Ukraine and strengthened security measures at the border. The aid is now in jeopardy.",
		Link:    "",
	},

	{
		Title:   "With Demise of Border Deal, No Clear Path for Ukraine and Israel Aid in Congress",
		Content: "The Senate’s $118.3 billion border deal that included aid to Ukraine appeared dead.",
		Link:    "",
	},

	{
		Title:   "Federal Appeals Court Rejects Trump’s Claim of Absolute Immunity",
		Content: "Former President Donald J. Trump is expected to continue his appeal to the Supreme Court.",
		Link:    "",
	},

	{
		Title:   "Judge in Trump’s Civil Fraud Case Asks If Key Witness Committed Perjury",
		Content: "Allen H. Weisselberg, former chief financial officer of the Trump Organization, has already done one stint in Rikers. ",
		Link:    "",
	},

	{
		Title:   "Jennifer Crumbley, Mother of Michigan School Gunman, Guilty of Manslaughter",
		Content: "Jennifer Crumbley, center, entered the courtroom in Oakland County, Mich., on Monday.",
		Link:    "",
	},

	{
		Title:   "What Israeli Soldiers’ Social Media Videos in Gaza Reveal",
		Content: "An analysis of social media videos found Israeli soldiers filming themselves in Gaza and destroying what appears to be civilian property. The footage provides a rare and unsanctioned window into the war.",
		Link:    "",
	},

	{
		Title:   "Internal Israeli Report Says a Fifth of the Remaining Hostages Have Died",
		Content: "Photographs of hostages seized on Oct. 7, on a wall in Tel Aviv on Monday.",
		Link:    "",
	},

	{
		Title:   "Nevada’s Weird Primary and Caucuses: What to Watch For",
		Content: "Former President Donald J. Trump is expected to receive all 26 delegates at stake in the Nevada contest, with nobody contesting him in Thursday’s caucuses.",
		Link:    "",
	},

	{
		Title:   "The Most Surprising Ways 2024 Candidates Spent Donor Money",
		Content: "Gov. Ron DeSantis of Florida serving ice cream in August in Iowa. Candidates and political committees have spent more than $10,000 on ice cream events so far this cycle.",
		Link:    "",
	},

	{
		Title:   "Where Has Tracy Chapman Been? Her Grammys Triumph Has Fans Wondering",
		Content: "Tracy Chapman made a rare public appearance to perform her 1988 song “Fast Car” with the country singer Luke Combs at the Grammy Awards on Sunday night.",
		Link:    "",
	},

	{
		Title:   "King Charles’s Cancer Diagnosis May Reshape How U.K. Monarchy Works",
		Content: "King Charles III meeting midwives and nurses at Buckingham Palace last year. Public engagements have long been a key image-maker for Britain’s royal family.",
		Link:    "",
	},

	{
		Title:   "Tucker Carlson Says He Will Soon Interview Putin",
		Content: "Tucker Carlson last year. An interview would come at a critical time for the war in Ukraine, with American aid to Kyiv stalled in Congress. ",
		Link:    "",
	},

	{
		Title:   "Alaska Airlines 737 May Have Left Boeing Factory Missing Bolts, N.T.S.B. Says",
		Content: "Boeing has been scrambling to contain the fallout after a panel on this Alaska Airlines plane blew out midair, terrifying passengers and forcing an emergency landing.",
		Link:    "",
	},

	{
		Title:   "Endless Range, Boundless Swagger: Why Caitlin Clark Is Different",
		Content: "At home and on the road, Caitlin Clark and Iowa have played before raucous, full-capacity crowds.",
		Link:    "",
	},

	{
		Title:   "In Poland, I Saw What a Second Trump Term Could Do to America",
		Content: "Being in Warsaw was inspiring and sobering.",
		Link:    "",
	},

	{
		Title:   "There’s a Tax Season Villain, and It’s Not the I.R.S.",
		Content: "For decades, American presidents have promised free and easy tax filling. Why have they consistently failed to deliver?",
		Link:    "",
	},

	{
		Title:   "The Immunity Ruling Is a Blow to Trump’s Monarchy",
		Content: "The appeals court ruling supports the American Revolution.",
		Link:    "",
	},

	{
		Title:   "Joe Manchin Would Like Your Attention, and He is Not Alone",
		Content: "Let the third-party follies begin.",
		Link:    "",
	},

	{
		Title:   "Peace in Ukraine",
		Content: "The Donetsk region last month.",
		Link:    "",
	},

	{
		Title:   "Federal Records Show Increasing Use of Solitary Confinement for Immigrants",
		Content: "Pastor Steven Tendo was detained by ICE for 26 months, including recurring stints in solitary confinement.",
		Link:    "",
	},

	{
		Title:   "Toby Keith, Larger-Than-Life Country Music Star, Dies at 62",
		Content: "Toby Keith performing in Nashville in 2018. His songs ranged from traditional honky-tonk to pop-country balladry and Southern rock. ",
		Link:    "",
	},

	{
		Title:   "Scores of NYC Public Housing Workers Charged in Record Corruption Case",
		Content: "Damian Williams, the U.S. attorney for the Southern District of New York, said at a news conference that nearly one-third of the New York City Housing Authorities buildings were involved in the bribery scheme.",
		Link:    "",
	},

	{
		Title:   "Southern California Braces for Potential Mudslides as Rain Continues",
		Content: "A collapsed structure in a Los Angeles residential neighborhood on Monday.",
		Link:    "",
	},

	{
		Title:   "How One Couple Modernized Their 19th-Century Salem Home",
		Content: "Here’s how one couple brought their Federal house in Salem, Mass., into the 21st century — with ‘color, character and eccentricity.’",
		Link:    "",
	},

	{
		Title:   "Seizures of Psychedelic Mushrooms Rise in U.S. as Demand Grows",
		Content: "Psychoactive mushrooms being grown in Oregon. ",
		Link:    "",
	},

	{
		Title:   "Endless Range, Boundless Swagger: Why Caitlin Clark Is Different",
		Content: "At home and on the road, Caitlin Clark and Iowa have played before raucous, full-capacity crowds.",
		Link:    "",
	},

	{
		Title:   "Southern California Braces for Potential Mudslides as Rain Continues",
		Content: "Weather experts warned that additional rain on top of saturated soil in the Los Angeles region could still cause hillsides to collapse.",
		Link:    "",
	},

	{
		Title:   "From Ferguson to Gaza: How African Americans Bonded With Palestinian Activists",
		Content: "Sandra Tamari, a Palestinian American, mourned the deaths of Michael Brown, killed in Ferguson, and Muhammad Abu Khdeir, killed in Jerusalem. She saw them both as teenagers stolen from their families by racially motivated violence, and joined protests in Ferguson.",
		Link:    "",
	},

	{
		Title:   "Clyde Taylor, Literary Scholar Who Elevated Black Cinema, Dies at 92",
		Content: "Clyde Taylor in the 1970s, when he was at the epicenter of a push to bring the study of Black culture into academia.",
		Link:    "",
	},

	{
		Title:   "Texas Company Was Behind Voter Robocalls That Impersonated Biden, N.H. Says",
		Content: "President Biden during a campaign event in Las Vegas on Sunday. Robocalls impersonating Mr. Biden were made to New Hampshire voters ahead of the state’s primary last month.",
		Link:    "",
	},

	{
		Title:   "Jennifer Crumbley, Mother of Michigan School Gunman, Guilty of Manslaughter",
		Content: "Jennifer Crumbley, center, entered the courtroom in Oakland County, Mich., on Monday.",
		Link:    "",
	},

	{
		Title:   "On the Border, Republicans Set a Trap, Then Fell Into It",
		Content: "Republicans are rapidly abandoning a legislative compromise that would have given aid to Ukraine and strengthened security measures at the border. The aid is now in jeopardy.",
		Link:    "",
	},

	{
		Title:   "Trump Says Ronna McDaniel ‘Knows’ She Should Step Down as R.N.C. Chair",
		Content: "Ronna McDaniel, chairwoman of the Republican Party, speaking at a Youth Advisory Council Roundtable in Manchester, N.H., last month.",
		Link:    "",
	},

	{
		Title:   "With Demise of Border Deal, No Clear Path for Ukraine and Israel Aid in Congress",
		Content: "The Senate’s $118.3 billion border deal that included aid to Ukraine appeared dead.",
		Link:    "",
	},

	{
		Title:   "The Most Surprising Ways 2024 Candidates Spent Donor Money",
		Content: "Gov. Ron DeSantis of Florida serving ice cream in August in Iowa. Candidates and political committees have spent more than $10,000 on ice cream events so far this cycle.",
		Link:    "",
	},

	{
		Title:   "Federal Appeals Court Rejects Trump’s Claim of Absolute Immunity",
		Content: "Former President Donald J. Trump is expected to continue his appeal to the Supreme Court.",
		Link:    "",
	},

	{
		Title:   "Explaining a Major Education Settlement in California",
		Content: "A Los Angeles Unified School District student attended an online class in 2020.",
		Link:    "",
	},

	{
		Title:   "FAA Chief Pledges ‘More Boots on the Ground’ to Monitor Boeing",
		Content: "Mike Whitaker, the Federal Aviation Administration’s top official, has taken a hard line with Boeing after a door panel blew out of a 737 Max 9 jet last month.",
		Link:    "",
	},

	{
		Title:   "Nevada’s Weird Primary and Caucuses: What to Watch For",
		Content: "Former President Donald J. Trump is expected to receive all 26 delegates at stake in the Nevada contest, with nobody contesting him in Thursday’s caucuses.",
		Link:    "",
	},

	{
		Title:   "Group in Michigan Urges Protest Vote Against Biden Over Israel-Gaza War",
		Content: "President Biden speaking to autoworkers in Michigan, which will be a key battleground state in November.",
		Link:    "",
	},

	{
		Title:   "Republican-Led House to Vote on Alejandro Mayorkas Impeachment",
		Content: "Alejandro N. Mayorkas, the homeland security secretary, would become the only sitting cabinet member to be impeached in American history.",
		Link:    "",
	},

	{
		Title:   "When Mental Health Treatment Becomes a Political Identity",
		Content: "County Judge Lina Hidalgo of Harris County in Houston, the third-most populous county in the country.",
		Link:    "",
	},

	{
		Title:   "Biden Calls Los Angeles Mayor and Promises Aid for Storm Damage",
		Content: "President Biden arrived at Joint Base Andrews, Md., on Monday, after campaign events.",
		Link:    "",
	},

	{
		Title:   "Aurora, Colo., Pays $1.9 Million to Black Family Wrongly Detained by Police",
		Content: "Brittney Gilliam and her family agreed to settle their federal civil rights lawsuit against Aurora, Colo, for $1.9 million.",
		Link:    "",
	},

	{
		Title:   "Nikki Haley Requests Secret Service Protection as She Faces Rising Threats",
		Content: "Nikki Haley during a campaign event in Charleston, S.C., on Sunday. Ms. Haley has been the target of at least two “swatting” hoax calls since December.",
		Link:    "",
	},

	{
		Title:   "Trump Couldn’t Shut Down the Border. Can Biden?",
		Content: "President Biden has not tested the limits of immigration law, unlike his immediate predecessor, Donald J. Trump.",
		Link:    "",
	},

	{
		Title:   "Biden Threatens to Veto Bill That Would Help Israel but Not Ukraine",
		Content: "President Biden departing Las Vegas on Monday. The Biden administration said the president would veto the House bill if it came to his desk.",
		Link:    "",
	},

	{
		Title:   "Fact-Checking Claims That Senate Bill Allows 5,000 Unauthorized Immigrants a Day",
		Content: "Under current law, a migrant who crosses the border illegally can apply for asylum, which the border deal could change.",
		Link:    "",
	},

	{
		Title:   "Mudslide From Storm Damages Homes in Studio City Section of Los Angeles",
		Content: "Lockridge Road in Studio City, Calif., was covered with mud and household debris on Monday.",
		Link:    "",
	},

	{
		Title:   "G.O.P. Backlash to Border Deal Reflects Vanishing Ground for a Compromise",
		Content: "Senator James Lankford, the lead Republican negotiator on the border deal, said that members of his party who complained that they needed more time to read through the bill were rushing to denounce it on social media.",
		Link:    "",
	},

	{
		Title:   "U.S. Strikes Hit Most of Targets in Iraq and Syria, Pentagon Says",
		Content: "Maj. Gen. Pat Ryder, the Pentagon spokesman, speaking to reporters on Monday.",
		Link:    "",
	},

	{
		Title:   "How Trump Uses the Power and Imagery of His Presidency",
		Content: "Former President Donald J. Trump leaves a campaign rally while surrounded by Secret Service agents in Durham, N.H., in December.",
		Link:    "",
	},

	{
		Title:   "The Retribution Presidency",
		Content: "Supporters lining up to see Former President Donald J. Trump speak in Concord, N.H.",
		Link:    "",
	},

	{
		Title:   "Ex-Day Care Director Who Fed Children Melatonin Gets 6 Months in Jail",
		Content: "Tonya Rachelle Voris, the former director of Kidz Life Childcare Ministry in Cumberland, Ind., was sentenced on Friday to six months in jail.",
		Link:    "",
	},

	{
		Title:   "Larry David Breaks Georgia’s Voting Law in ‘Curb Your Enthusiasm’",
		Content: "Comedy’s crankiest star usually breaks social norms. This time, he ran afoul of voting legislation passed in Georgia in 2021.",
		Link:    "",
	},

	{
		Title:   "Biden Meets With Culinary Workers on Eve of Nevada Primary",
		Content: "President Biden met with members of the Culinary Workers Union in Las Vegas on Monday.",
		Link:    "",
	},

	{
		Title:   "On the Border, Republicans Set a Trap, Then Fell Into It",
		Content: "Republicans are rapidly abandoning a legislative compromise that would have given aid to Ukraine and strengthened security measures at the border. The aid is now in jeopardy.",
		Link:    "",
	},

	{
		Title:   "Federal Appeals Court Rejects Trump’s Claim of Absolute Immunity",
		Content: "Former President Donald J. Trump is expected to continue his appeal to the Supreme Court.",
		Link:    "",
	},

	{
		Title:   "The Trump Election Immunity Ruling, Annotated",
		Content: "A three-judge appeals court panel rejected former President Donald J. Trump’s claim that he was immune to criminal charges. Read its decision, with Times analysis.",
		Link:    "",
	},

	{
		Title:   "What to Watch in Nevada’s Weird Election Week",
		Content: "Former President Donald J. Trump is expected to receive all 26 delegates at stake in the Nevada contest, with nobody contesting him in Thursday’s caucuses.",
		Link:    "",
	},

	{
		Title:   "Alaska Airlines 737 May Have Left Boeing Factory Missing Bolts, N.T.S.B. Says",
		Content: "Boeing has been scrambling to contain the fallout after a panel on this Alaska Airlines plane blew out midair, terrifying passengers and forcing an emergency landing.",
		Link:    "",
	},

	{
		Title:   "Texas Company Was Behind Voter Robocalls That Impersonated Biden, N.H. Says",
		Content: "President Biden during a campaign event in Las Vegas on Sunday. Robocalls impersonating Mr. Biden were made to New Hampshire voters ahead of the state’s primary last month.",
		Link:    "",
	},

	{
		Title:   "Trump Says Ronna McDaniel ‘Knows’ She Should Step Down as R.N.C. Chair",
		Content: "Ronna McDaniel, chairwoman of the Republican Party, speaking at a Youth Advisory Council Roundtable in Manchester, N.H., last month.",
		Link:    "",
	},

	{
		Title:   "With Demise of Border Deal, No Clear Path for Ukraine and Israel Aid in Congress",
		Content: "The Senate’s $118.3 billion border deal that included aid to Ukraine appeared dead.",
		Link:    "",
	},

	{
		Title:   "The Most Surprising Ways 2024 Candidates Spent Donor Money",
		Content: "Gov. Ron DeSantis of Florida serving ice cream in August in Iowa. Candidates and political committees have spent more than $10,000 on ice cream events so far this cycle.",
		Link:    "",
	},

	{
		Title:   "Federal Records Show Increasing Use of Solitary Confinement for Immigrants",
		Content: "Pastor Steven Tendo was detained by ICE for 26 months, including recurring stints in solitary confinement.",
		Link:    "",
	},

	{
		Title:   "U.S. Ambassadors in the Pacific Urge Action on Ukraine, Israel and Border Bill",
		Content: "A demonstration in support of Palestinians in Gaza, in Melbourne, Australia, in November.",
		Link:    "",
	},

	{
		Title:   "FAA Chief Pledges ‘More Boots on the Ground’ to Monitor Boeing",
		Content: "Mike Whitaker, the Federal Aviation Administration’s top official, has taken a hard line with Boeing after a door panel blew out of a 737 Max 9 jet last month.",
		Link:    "",
	},

	{
		Title:   "Group in Michigan Urges Protest Vote Against Biden Over Israel-Gaza War",
		Content: "President Biden speaking to autoworkers in Michigan, which will be a key battleground state in November.",
		Link:    "",
	},

	{
		Title:   "Republican-Led House to Vote on Alejandro Mayorkas Impeachment",
		Content: "Alejandro N. Mayorkas, the homeland security secretary, would become the only sitting cabinet member to be impeached in American history.",
		Link:    "",
	},

	{
		Title:   "Biden Has Openings for a Comeback on Two Weak Points",
		Content: "President Biden has good economic news to promote and a possible border tactic to deploy. ",
		Link:    "",
	},

	{
		Title:   "Yellen Says Stable Financial System Is Key to U.S. Economic Strength",
		Content: "In House testimony on Tuesday, Treasury Secretary Janet Yellen is expected to face questions about what has been done to safeguard the financial system since Silicon Valley Bank and Signature Bank collapsed.",
		Link:    "",
	},

	{
		Title:   "Nikki Haley Requests Secret Service Protection as She Faces Rising Threats",
		Content: "Nikki Haley during a campaign event in Charleston, S.C., on Sunday. Ms. Haley has been the target of at least two “swatting” hoax calls since December.",
		Link:    "",
	},

	{
		Title:   "Trump Couldn’t Shut Down the Border. Can Biden?",
		Content: "President Biden has not tested the limits of immigration law, unlike his immediate predecessor, Donald J. Trump.",
		Link:    "",
	},

	{
		Title:   "Biden Threatens to Veto Bill That Would Help Israel but Not Ukraine",
		Content: "President Biden departing Las Vegas on Monday. The Biden administration said the president would veto the House bill if it came to his desk.",
		Link:    "",
	},

	{
		Title:   "Fact-Checking Claims That Senate Bill Allows 5,000 Unauthorized Immigrants a Day",
		Content: "Under current law, a migrant who crosses the border illegally can apply for asylum, which the border deal could change.",
		Link:    "",
	},

	{
		Title:   "Blinken Meets Saudi Crown Prince on Mideast Push for Pause in Gaza War",
		Content: "Secretary of State Antony J. Blinken meeting with Crown Prince Mohammed bin Salman of Saudi Arabia in Riyadh on Monday.",
		Link:    "",
	},

	{
		Title:   "G.O.P. Backlash to Border Deal Reflects Vanishing Ground for a Compromise",
		Content: "Senator James Lankford, the lead Republican negotiator on the border deal, said that members of his party who complained that they needed more time to read through the bill were rushing to denounce it on social media.",
		Link:    "",
	},

	{
		Title:   "U.S. Strikes Hit Most of Targets in Iraq and Syria, Pentagon Says",
		Content: "Maj. Gen. Pat Ryder, the Pentagon spokesman, speaking to reporters on Monday.",
		Link:    "",
	},

	{
		Title:   "How Trump Uses the Power and Imagery of His Presidency",
		Content: "Former President Donald J. Trump leaves a campaign rally while surrounded by Secret Service agents in Durham, N.H., in December.",
		Link:    "",
	},

	{
		Title:   "The Retribution Presidency",
		Content: "Supporters lining up to see Former President Donald J. Trump speak in Concord, N.H.",
		Link:    "",
	},

	{
		Title:   "Larry David Breaks Georgia’s Voting Law in ‘Curb Your Enthusiasm’",
		Content: "Comedy’s crankiest star usually breaks social norms. This time, he ran afoul of voting legislation passed in Georgia in 2021.",
		Link:    "",
	},

	{
		Title:   "Biden Meets With Culinary Workers on Eve of Nevada Primary",
		Content: "President Biden met with members of the Culinary Workers Union in Las Vegas on Monday.",
		Link:    "",
	},

	{
		Title:   "Plunge in New York Community Bank’s Stock Stirs Fears of Wider Crisis",
		Content: "An Atlantic Bank branch, part of New York Community Bank, in New York last month. The parent company’s stock plummeted after it released an ugly earnings report last week.",
		Link:    "",
	},

	{
		Title:   "$7 Million for 30 Seconds? To Advertisers, the Super Bowl Is Worth It.",
		Content: "Hellmann’s Mayonnaise is one of the companies that will be advertising during Sunday’s Super Bowl.",
		Link:    "",
	},

	{
		Title:   "Meta Calls for Industry Effort to Label A.I.-Generated Content",
		Content: "Nick Clegg, president of global affairs and communications at Meta, is pushing other companies to help identify A.I.-created content.",
		Link:    "",
	},

	{
		Title:   "FAA Chief Pledges ‘More Boots on the Ground’ to Monitor Boeing",
		Content: "Mike Whitaker, the Federal Aviation Administration’s top official, has taken a hard line with Boeing after a door panel blew out of a 737 Max 9 jet last month.",
		Link:    "",
	},

	{
		Title:   "Yellen Says Stable Financial System Is Key to U.S. Economic Strength",
		Content: "In House testimony on Tuesday, Treasury Secretary Janet Yellen is expected to face questions about what has been done to safeguard the financial system since Silicon Valley Bank and Signature Bank collapsed.",
		Link:    "",
	},

	{
		Title:   "Summer Has Long Stressed Electric Grids. Now Winter Does, Too.",
		Content: "Star Pizza lost power in its 50-year-old location in Houston during a winter storm in 2021.",
		Link:    "",
	},

	{
		Title:   "Boeing, Still Recovering From Max 8 Crashes, Faces a New Crisis",
		Content: "Boeing’s production facility in Renton, Wash. Before two deadly plane crashes involving the company’s 737 Max 8 aircraft five years ago, Boeing was producing around 52 planes per month.",
		Link:    "",
	},

	{
		Title:   "Why Is Big Tech Still Cutting Jobs?",
		Content: "Microsoft made job cuts in its video game division in January.",
		Link:    "",
	},

	{
		Title:   "Alaska Airlines 737 May Have Left Boeing Factory Missing Bolts, N.T.S.B. Says",
		Content: "Boeing has been scrambling to contain the fallout after a panel on this Alaska Airlines plane blew out midair, terrifying passengers and forcing an emergency landing.",
		Link:    "",
	},

	{
		Title:   "Could This Be the Most Expensive Super Bowl Ever?",
		Content: "Including the cost of flights, lodging, assorted travel expenses and entry into Allegiant Stadium, Super Bowl LVIII may be the most expensive to attend. ",
		Link:    "",
	},

	{
		Title:   "President of Powerful Service Workers Union Will Step Down",
		Content: "Mary Kay Henry, the president of the Service Employees International Union since 2010, said she would not run for another term.",
		Link:    "",
	},

	{
		Title:   "BP’s Shares Rise After It Says It Will Increase Oil Output",
		Content: "A BP gas station in New York. The company’s chief executive said that investments in oil and gas would continue.",
		Link:    "",
	},

	{
		Title:   "Adam Neumann Wants to Take Over WeWork",
		Content: "Lawyers for Adam Neumann accused WeWork of stonewalling his takeover approach.",
		Link:    "",
	},

	{
		Title:   "Dartmouth Players Are Employees Who Can Unionize, U.S. Official Says",
		Content: "Robert McRae III, left, and Romeo Myrthil of the Dartmouth men’s basketball team, which is seeking to join the Service Employees International Union.",
		Link:    "",
	},

	{
		Title:   "Banks Sue Regulators Over Anti-Redlining Rule",
		Content: "A construction barrier outside the Federal Reserve building last spring. The Fed, the F.D.I.C. and the Office of the Comptroller of the Currency were sued by banking groups on Monday.",
		Link:    "",
	},

	{
		Title:   "Grammy Audience Jumps to 16.9 Million",
		Content: "From left, Phoebe Bridgers, Lucy Dacus and Taylor Swift celebrated victories at the Grammy Awards on Sunday night.",
		Link:    "",
	},

	{
		Title:   "Facing Another Boeing Crisis, the FAA Takes a Harder Line",
		Content: "The Federal Aviation Administration grounded scores of Boeing 737 Max 9 jets in January after a door panel blew out of an Alaska Airlines flight in midair.",
		Link:    "",
	},

	{
		Title:   "The Hospitality Group Redefining Korean Restaurants in New York",
		Content: "The scene at Ariari, one of 21 restaurants in New York owned or co-owned by Hand Hospitality. The company has helped turn the city into a hub for innovative Korean dining.",
		Link:    "",
	},

	{
		Title:   "Taylor Swift Prop Bets Dominate Super Bowl Action",
		Content: "Taylor Swift has attended a dozen Kansas City Chiefs games this season, giving oddsmakers plenty of material to work with when building bets around her expected appearance at the Super Bowl.",
		Link:    "",
	},

	{
		Title:   "CNN’s New Morning Strategy: More News, Less Banter",
		Content: "The CNN chairman, Mark Thompson, is shaking up the network’s morning lineup in his first significant programming move.",
		Link:    "",
	},

	{
		Title:   "Yandex Reaches $5 Billion Deal to Exit Russia",
		Content: "Among the assets that Yandex will sell are a popular internet browser and Russia’s main food delivery and taxi-hailing apps.",
		Link:    "",
	},

	{
		Title:   "Snap Lays Off 10% of Its Work Force",
		Content: "Snap is expected to report its fourth-quarter earnings on Tuesday.",
		Link:    "",
	},

	{
		Title:   "The Jobs Conundrum: Questions About Wages Persist",
		Content: "Jobs seem plentiful, but a large group of voters are feeling downbeat about inflation and the economy.",
		Link:    "",
	},

	{
		Title:   "Top U.S. Treasury Officials to Visit Beijing for Economic Talks",
		Content: "The visit could pave the way for a second trip to China by Treasury Secretary Janet L. Yellen. She traveled to Beijing last summer.",
		Link:    "",
	},

	{
		Title:   "Fake and Explicit Images of Taylor Swift Started on 4chan, Study Says",
		Content: "Graphika found a thread of messages on 4chan that encouraged people to try to evade safeguards set up by image generator tools to create the recent offensive deepfake images of Taylor Swift.",
		Link:    "",
	},

	{
		Title:   "Big Companies Cashed In on Mississippi’s Water. Small Towns Paid the Price.",
		Content: "Defunct water meters at City Hall in Moss Point, Miss.",
		Link:    "",
	},

	{
		Title:   "How Nevada Is Pushing to Generate Jobs Beyond the Casinos",
		Content: "Reno, Nev., which has long been in the shadow of Las Vegas, is trying to reinvent itself.",
		Link:    "",
	},

	{
		Title:   "Why Are Americans Wary While the Economy Is Healthy? Look at Nevada.",
		Content: "Economic shocks over two decades, combined with reliance on volatile casinos, have undermined confidence despite an economy that’s bustling.",
		Link:    "",
	},

	{
		Title:   "Boeing Finds More Problems With 737 Max, Risking Delivery Delays",
		Content: "A Boeing 737 Max 9 jet under construction at the company’s factory in Renton, Wash, in 2017.",
		Link:    "",
	},

	{
		Title:   "Samsung’s Lee Jae-yong Acquitted in Stock, Accounting Fraud Case",
		Content: "Lee Jae-yong, the chairman of Samsung Electronics, arriving at the Seoul Central District Court in Seoul on Monday.",
		Link:    "",
	},

	{
		Title:   "What Business Leaders Are Saying About the Red Sea Attacks",
		Content: "A Greek-owned ship at the Suez Shipyard Company in Egypt last month getting repairs on the damage caused by a Houthi missile attack in the Red Sea.",
		Link:    "",
	},

	{
		Title:   "What’s Behind This $10 Chicken Over Rice? An $18,000 Permit.",
		Content: "Blame rising costs, shrunken crowds and a black-market permit trade. Chicken over rice costs $10 at the Halal Plates, a cart in Lower Manhattan, up from $6 prepandemic.",
		Link:    "",
	},

	{
		Title:   "Fed Chair Powell Says Officials Need More ‘Good’ Data Before Cutting Rates",
		Content: "The Federal Reserve chair Jerome H. Powell in December. “We think we can be careful in approaching this decision just because of the strength that we’re seeing in the economy,” he said on “60 Minutes” on Sunday.",
		Link:    "",
	},

	{
		Title:   "Coping With the Loss of a Pet? They Can Offer Grief Support.",
		Content: "The death of a pet can leave its owner with complex feelings that they may struggle to process alone. “People sometimes think they’re going crazy,” said Colleen Rolland, the president of the Association for Pet Loss and Bereavement.",
		Link:    "",
	},

	{
		Title:   "Oscar Goodman, Former Las Vegas Mayor, Won’t Be at the Super Bowl",
		Content: "Oscar Goodman, in his Las Vegas home, was the city’s mayor from 1999 to 2011. ",
		Link:    "",
	},

	{
		Title:   "N.F.L.’s Rapid Embrace of Gambling Creates Mixed Signals",
		Content: "The BetMGM Sportsbook & Lounge in Maryland. The National Football League is part of a growing apparatus that encourages casual fans to regularly place wagers on games.",
		Link:    "",
	},

	{
		Title:   "What’s My D.E.I. Training? My Own Life.",
		Content: "Sometimes brutal honesty isn’t the best policy when interviewing for a new job.",
		Link:    "",
	},

	{
		Title:   "Medicare Recipients Can Save on Drugs in 2024",
		Content: "Federal law set caps on some spending for seniors who are enrolled in Part D prescription drug plans. Those paying for high-cost medicines will benefit.",
		Link:    "",
	},

	{
		Title:   "Biden to Sit Out Super Bowl Interview",
		Content: "President Biden at his campaign headquarters in Delaware on Saturday. In a tradition dating to 2009, presidents typically record an interview with the network broadcasting the Super Bowl.",
		Link:    "",
	},

	{
		Title:   "Target Pulls Magnet Kit That Misidentified Three Black Leaders",
		Content: "Target said it would no longer sell the magnet kit online or in its stores, and that it had “ensured the product’s publisher is aware of the errors.”",
		Link:    "",
	},

	{
		Title:   "What Happens to FTX Clawback Cases if the Company Repays Its Creditors?",
		Content: "The FTX founder Sam Bankman-Fried was convicted of seven charges of fraud and conspiracy in November.",
		Link:    "",
	},

	{
		Title:   "What Should Boeing Do to Fix Its Longstanding Problems?",
		Content: "A panel blew off a 737 Max 9 plane during an Alaska Airlines flight this month.",
		Link:    "",
	},

	{
		Title:   "Turkey’s Central Bank Chief Steps Down Amid Long Inflation Battle",
		Content: "Hafize Gaye Erkan was Turkey’s fifth central bank chief in five years.",
		Link:    "",
	},

	{
		Title:   "How Much for Thin Mints? Some Girl Scouts Raise Cookie Prices.",
		Content: "Caramel deLites and Thin Mints Girl Scout cookies.",
		Link:    "",
	},

	{
		Title:   "A City Built on Steel Tries to Reverse Its Decline",
		Content: "Gary, Ind., home of U.S. Steel’s largest mill, is seeking a fresh economic start under a new mayor determined to draw new businesses and residents.",
		Link:    "",
	},

	{
		Title:   "Anxiety, Mood Swings and Sleepless Nights: Life Near a Bitcoin Mine",
		Content: "Pushed by an advocacy group, Arkansas became the first state to shield noisy cryptocurrency operators from unhappy neighbors. A furious backlash has some lawmakers considering a statewide ban.",
		Link:    "",
	},

	{
		Title:   "U.S. Hits Back at Iran With Sanctions, Criminal Charges and Airstrikes",
		Content: "A photograph released by Iranian state news in August shows members of the Islamic Revolutionary Guards Corps.",
		Link:    "",
	},

	{
		Title:   "Joe Madison, Radio Host and Civil Rights Activist, Dies at 74",
		Content: "Joe Madison in 2015. His activism included a hunger strike and handcuffing himself to the Sudanese Embassy in Washington.",
		Link:    "",
	},

	{
		Title:   "U.S. Leading Soft Landing for Global Economy",
		Content: "Americans have only slowly spent down the savings they amassed during the early pandemic years, so the money has continued to trickle through the economy like a slow-release booster shot.",
		Link:    "",
	},

	{
		Title:   "Meta’s Stock Surges After Jump in Profits",
		Content: "The trillion-dollar social media giant’s market value soared after it reported a robust profits driven by digital advertising.",
		Link:    "",
	},

	{
		Title:   "He’s Lost His Marriage, His Followers and His Lamborghini",
		Content: "Ben Armstrong, better known as BitBoy, was once the most popular cryptocurrency YouTuber in the world. Now his empire has collapsed.",
		Link:    "",
	},

	{
		Title:   "For Biden, a Sunny Economy Could Finally Be a Potential Gain",
		Content: "President Biden has struggled to sell voters on the positive signs in the economy under his watch.",
		Link:    "",
	},

	{
		Title:   "6 Reasons That It’s Hard to Get Your Wegovy and Other Weight-Loss Prescriptions",
		Content: "An array of obstacles makes it difficult for patients to obtain Wegovy or Zepbound. Finding Wegovy is “like winning the lottery,” one nurse practitioner said.",
		Link:    "",
	},

	{
		Title:   "U.S. Job Gains Countered Economists’ Expectations for a Hiring Slowdown",
		Content: "Jobs and wage gains remain robust despite the Federal Reserve’s campaign to cool the economy, surprising policymakers and economic forecasters.",
		Link:    "",
	},

	{
		Title:   "Blockbuster Jobs Report Backs Up Fed’s Patience as It Waits to Cut Rates",
		Content: "Given continued strength of the economy, the Federal Reserve is unlikely to feel pressure to cut interest rates at its next meeting in March.",
		Link:    "",
	},

	{
		Title:   "F.T.C. Warns Dozens of Funeral Homes to Provide Accurate Costs to Callers",
		Content: "The agency said an “undercover phone sweep” of more than 250 homes found that 38 failed to provide prices or supplied inconsistent prices in separate calls.",
		Link:    "",
	},
}
