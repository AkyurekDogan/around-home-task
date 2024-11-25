CREATE TABLE public.partner (
	id varchar(10) NOT NULL,
	"name" varchar(250) NOT NULL,
	"location" public.geography(point, 4326) NOT NULL,
	radius int4 NOT NULL,
	CONSTRAINT partner_pk PRIMARY KEY (id)
);

CREATE TABLE public.rating (
	partner_id varchar(10) NOT NULL,
	avg int4 NOT NULL
);

ALTER TABLE public.rating ADD CONSTRAINT rating_partner_fk FOREIGN KEY (partner_id) REFERENCES public.partner(id);


CREATE TABLE public.skill (
	partner_id varchar(10) NOT NULL,
	craftsmanship_tags _text NOT NULL
);

ALTER TABLE public.skill ADD CONSTRAINT skill_partner_fk FOREIGN KEY (partner_id) REFERENCES public.partner(id);

insert into public.partner (id,"name","location",radius) 
values ('px12xxx123','X Engineering - Gesundrunen/Berlin', ST_SetSRID(ST_MakePoint(52.550385, 13.380968), 4326), 20),
       ('py12yyy123','Y Engineering - Spandau/Berlin',ST_SetSRID(ST_MakePoint(52.537976, 13.204127), 4326), 30),
       ('pz12zzz123','Z Engineering - Hauptbahnhof/Berlin',ST_SetSRID(ST_MakePoint(52.526804, 13.365678), 4326), 20),
       ('pt12ttt123','T Engineering - Schoneberg/Berlin',ST_SetSRID(ST_MakePoint(52.490920, 13.359833), 4326), 10),
       ('pk12kkk123','K Engineering - Neukoln/Berlin',ST_SetSRID(ST_MakePoint(52.475945, 13.446991), 4326), 30),
       ('pl12lll123','L Engineering - Prenzlauer/Berlin',ST_SetSRID(ST_MakePoint(52.540394, 13.423423), 4326), 30),
       ('pm12mmm123','M Engineering - Wittenau/Berlin',ST_SetSRID(ST_MakePoint(52.601897, 13.334707), 4326), 30),
       ('pn12nnn123','N Engineering - Adlershof/Berlin',ST_SetSRID(ST_MakePoint(52.437278, 13.534422), 4326), 30),
       ('po12ooo123','O Engineering - Lichtenberg/Berlin',ST_SetSRID(ST_MakePoint(52.535544, 13.497926), 4326), 30),
       ('pp12ppp123','P Engineering - Rodow/Berlin',ST_SetSRID(ST_MakePoint(52.422061, 13.495607), 4326),30),
       ('pa12aaa123','A Engineering - Potsdam/Brendenburg',ST_SetSRID(ST_MakePoint(52.404744, 13.045431), 4326), 40),
       ('pb12bbb123','B Engineering - Falkensee/Brendenburg',ST_SetSRID(ST_MakePoint(52.560051, 13.078569), 4326), 40),
       ('pc12ccc123','C Engineering - Hennigsdorf/Brendenburg',ST_SetSRID(ST_MakePoint(52.645989, 13.199716), 4326), 40),
       ('pd12ddd123','D Engineering - Bernau/Brendenburg',ST_SetSRID(ST_MakePoint(52.679390, 13.581343), 4326), 40),
       ('pe12eee123','E Engineering - Wegendorf/Brendenburg',ST_SetSRID(ST_MakePoint(52.601588, 13.750378), 4326), 70),
       ('pf12fff123','F Engineering - Magdeburg/Brendenburg',ST_SetSRID(ST_MakePoint(52.151980, 11.617793), 4326), 70),
       ('pg12ggg123','G Engineering - Wittschock/Brendenburg',ST_SetSRID(ST_MakePoint(53.191976, 12.493398), 4326), 80),
       ('ph12hhh123','H Engineering - Ueckermunde/Brendenburg',ST_SetSRID(ST_MakePoint(53.724077, 14.024921), 4326), 120),
       ('pi12iii123','I Engineering - Schorfheide/Brendenburg',ST_SetSRID(ST_MakePoint(52.865031, 13.748365), 4326), 30),
       ('pj12jjj123','J Engineering - Frankfurt Oder/Brendenburg',ST_SetSRID(ST_MakePoint(52.308907, 14.519605), 4326), 70);
       

insert into public.skill (partner_id,craftsmanship_tags)
values ('px12xxx123',ARRAY['wood', 'carpet', 'tiles']),
	   ('py12yyy123',ARRAY['wood', 'tiles']),
	   ('pz12zzz123',ARRAY['wood', 'tiles']),
	   ('pt12ttt123', ARRAY['wood', 'carpet', 'tiles']),
	   ('pk12kkk123', ARRAY['wood', 'carpet']),
	   ('pl12lll123',ARRAY['wood', 'carpet' ]),
	   ('pm12mmm123',ARRAY['carpet', 'tiles']),
	   ('pn12nnn123',ARRAY['wood', 'carpet', 'tiles']),
	   ('po12ooo123',ARRAY['wood', 'carpet' ]),
	   ('pp12ppp123', ARRAY['wood', 'carpet', 'tiles']),
	   ('pa12aaa123', ARRAY[ 'carpet' ]),
	   ('pb12bbb123',ARRAY['wood', 'carpet', 'tiles']),
	   ('pc12ccc123',ARRAY['wood', 'carpet', 'tiles']),
	   ('pd12ddd123',ARRAY[ 'carpet', 'tiles']),
	   ('pe12eee123',ARRAY[ 'carpet', 'tiles']),
	   ('pf12fff123',ARRAY['wood', 'carpet', 'tiles']),
	   ('pg12ggg123',ARRAY[ 'carpet', 'tiles']),
	   ('ph12hhh123',ARRAY['wood', 'carpet', 'tiles']),
	   ('pi12iii123',ARRAY[ 'carpet', 'tiles']),
	   ('pj12jjj123', ARRAY[ 'carpet', 'tiles']);
	   

insert into public.rating (partner_id,avg)
values ('px12xxx123',4),
	   ('py12yyy123',10),
	   ('pz12zzz123',7),
	   ('pt12ttt123', 8),
	   ('pk12kkk123', 5),
	   ('pl12lll123',9),
	   ('pm12mmm123',8),
	   ('pn12nnn123',7),
	   ('po12ooo123',9),
	   ('pp12ppp123', 6),
	   ('pa12aaa123', 6),
	   ('pb12bbb123',5),
	   ('pc12ccc123',4),
	   ('pd12ddd123',3),
	   ('pe12eee123',9),
	   ('pf12fff123',10),
	   ('pg12ggg123',7),
	   ('ph12hhh123',9),
	   ('pi12iii123',8),
	   ('pj12jjj123', 10);

CREATE INDEX idx_partner_location ON public.partner USING GIST (location);
