import type { Item, SlotItem, Team } from "./models";

const baseURL =
  "https://prd-game-a-granbluefantasy.akamaized.net/assets_en/img/sp/assets";
const lowBaseURl =
  "https://prd-game-a-granbluefantasy.akamaized.net/assets_en/img_low/sp/assets";

interface Config {
  item: Item;
  highRes: boolean;
  portrait: boolean;
  party: boolean;
  square: boolean;
}
export function getItemImageURL({
  item,
  highRes,
  portrait,
  party,
  square,
}: Config): string {
  let url = highRes ? baseURL : lowBaseURl;

  let shape = "m";
  if (square) shape = "s";

  switch (item.category) {
    case "character":
      // is Main Character
      if (item.element === "any" && item.type === "") {
        return portrait
          ? `${baseURL}/leader/raid_normal/${item.image_name}.jpg`
          : `${url}/leader/m/${item.id}_01.jpg`;
      }

      if (portrait) shape = "f";
      return `${url}/npc/${shape}/${item.image_name}.jpg`;

    case "summon":
      if (portrait) shape = "party_main";
      else if (party) shape = "party_sub";
      return `${url}/summon/${shape}/${item.image_name}.jpg`;

    case "weapon":
      if (portrait) shape = "ls";
      return `${url}/weapon/${shape}/${item.image_name}.jpg`;

    default:
      console.error("item has no category", item);
  }

  return "";
}

interface EmptyConfig {
  category: string;
  portrait: boolean;
}
export function getEmptyImagePath({ category, portrait }: EmptyConfig): string {
  switch (category) {
    case "character":
      return `/img/empty_character${portrait ? "_portrait" : ""}.jpg`;
    case "summon":
      return `/img/empty_summon${portrait ? "_portrait.png" : ".jpg"}`;
    case "weapon":
      return `/img/empty_weapon${portrait ? "_portrait.jpg" : ".png"}`;
    default:
      console.error("unknown category:", category);
  }

  return "";
}

export function getEmptyTeam(): Team {
  return {
    id: 0,
    name: "",
    url: "",
    character: new Array(7).fill(null),
    summon: new Array(8).fill(null),
    weapon: new Array(11).fill(null),
  };
}

export function sliceTeam(items: SlotItem[], start: number, end: number) {
  return items
    .slice(start, end)
    .map((item, i) =>
      item !== null ? item : { position: i + start, name: "-" },
    );
}

export function flattenItemsObject(items: Record<string, Item>) {
  return Object.values(items).map((i) => ({
    label: i.name.charAt(0).toUpperCase() + i.name.slice(1),
    id: i.id,
  }));
}

interface Identifiable {
  id: string;
}

export function areAllIn<A extends Identifiable, B extends Identifiable>(
  a: A[],
  b: (B | null)[],
): boolean {
  return a.every((itemA) =>
    b.some((itemB) => (itemB === null ? false : itemA.id === itemB.id)),
  );
}

export function isAnyIn<A extends Identifiable, B extends Identifiable>(
  a: A[],
  b: (B | null)[],
): boolean {
  return a.some((itemA) =>
    b.some((itemB) => (itemB === null ? false : itemA.id === itemB.id)),
  );
}

export function isIn<A extends Identifiable, B extends Identifiable>(
  checkAll: boolean,
  a: A[],
  b: (B | null)[],
): boolean {
  return checkAll ? areAllIn(a, b) : !isAnyIn(a, b);
}
